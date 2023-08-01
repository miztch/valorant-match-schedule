import datetime
import json
import logging
import os

import boto3
import constants
import database
import requests
import utils

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def calc_match_end_time(start_time, best_of):
    '''
    calculate
      - when the match ends
      - when the match info record(item in DynamoDB table) deleted
    '''
    # format timestamp for specification in Google Calendar
    start_time_dt = datetime.datetime.strptime(
        start_time, '%Y-%m-%dT%H:%M:%S%z')
    end_time_dt = start_time_dt + datetime.timedelta(hours=best_of)

    # set ttl as to delete match information from DynamoDB Table
    # 12 hours later than the match ends
    ttl_dt = end_time_dt + datetime.timedelta(hours=12)
    ttl = int(ttl_dt.timestamp())

    # return end time as string
    end_time_str = end_time_dt.strftime(
        '%Y-%m-%dT%H:%M:%S%z').replace('+0000', '+00:00')

    return end_time_str, ttl


def estimate_unspecified_region(event_name):
    '''
    fallback function for map_flag_to_region()
    return a region name estimated from event name or organizer name,
    or empty string if not matched any of patterns.
    '''
    logger.info('estimate region for event: {}'.format(event_name))
    sub_areas = constants.sub_areas
    organizers = constants.organizers

    region = ''

    # search from sub area name -> organizer name
    for dic in [sub_areas, organizers]:
        for region_name, strings in dic.items():
            for string in strings:
                if string in event_name:
                    region = region_name
                    logger.info('event: {} mapped to region: {}. matched word: {}'.format(
                        event_name, region, string))
                    break
        if region:
            break

    if not region:
        logger.warning(
            'event: {} was not mapped to any region'.format(event_name))

    return region


def map_flag_to_region(flag, region_map, event_name):
    '''
    get flag(usually it shows country) indicator and return region
    '''

    if flag in region_map:
        region = region_map[flag]
    elif flag == 'un':
        # 'un' flag (shows 'universal') is sometimes used in vlr.gg
        # usually for LATAM, MENA, APAC. Try to estimate from event name.
        logger.info(
            "flag 'un': needs fallback. event_name: {}".format(event_name))
        region = estimate_unspecified_region(event_name)
    else:
        logger.warning(
            "event: {} was not mapped to any region. flag: {}".format(event_name, flag))
        region = ''

    return region


def fetch_daily_matches(date):
    '''
    fetch match information for the specified day from api endpoint
    '''
    # configure datasource
    region_map = constants.countries
    international_events = constants.international_events

    domain = os.environ['API_DOMAIN_NAME']
    endpoint = 'https://' + domain + '/matches'

    logger.info('fetch matches list from: {}'.format(endpoint))

    # assemble request URI and header
    logger.info('fetch matches list for a day: {}'.format(date))
    uri = endpoint + '?date=' + date
    headers = constants.headers

    # due to api endpoint's behabior, requests.get.json() sometimes fails
    # retry till valid json can be gotten
    upcoming_matches = ''
    while True:
        utils.sleep()
        upcoming_matches = requests.get(uri, headers=headers)
        valid_json = utils.is_json(upcoming_matches.text)

        if valid_json:
            break

    matches = upcoming_matches.json()

    # pick up information for each individual match
    match_list = []
    for match in matches:
        # get region from flag indicator
        # skip if empty so that "Calendar Id" for Gcal cannot be determined
        flag = match['eventCountryFlag']
        event_name = utils.shorten(match['eventName'])
        region = map_flag_to_region(flag, region_map, event_name)

        if not region:
            continue

        match_id = match['id']
        teams = [team['title'] for team in match['teams']]

        # if international league match(ex. EMEA,Americas,Pacific)
        if event_name in international_events:
            region += '#INTERNATIONAL'

        # day x, upper/lower bracket, etc
        event_detail = utils.shorten(match['matchName'])

        start_time = match['startTime']
        best_of = match['bestOf'] if type(match['bestOf']) is int else '3'
        # calculate when the match ends (and when deleted from DynamoDB)
        end_time, ttl = calc_match_end_time(start_time, int(best_of))

        # match url
        match_uri = 'https://vlr.gg{}'.format(match['pagePath'])

        # assemble item as a dictionary
        item = {
            'match_id': match_id,
            'region': region,
            'team_home': teams[0],
            'team_away': teams[1],
            'event_name': event_name,
            'event_detail': event_detail,
            'best_of': best_of,
            'start_time': start_time,
            'end_time': end_time,
            'ttl': ttl,
            'match_uri': match_uri
        }
        # handle duplicated records sometimes included in API response
        # first come & first served
        if match_list and item['match_id'] in [item['match_id'] for item in match_list]:
            logger.info('match data (id: {}) duplicated, skipping. match info: {}'.format(
                item['match_id'], item))
        else:
            logger.info('add match info to the list: {}'.format(item))
            match_list.append(item)

    return match_list


def lambda_handler(event, context):
    records = event['Records']
    match_list = []

    for record in records:
        body = json.loads(record['body'])
        date = body['date']

        matches = fetch_daily_matches(date)
        match_list.extend(matches)

    database.insert(match_list)

    return {
        'matches_count': len(match_list)
    }
