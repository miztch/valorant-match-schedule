import datetime
import json
import logging
import os
import random
import time

import boto3
import requests

logger = logging.getLogger()
logger.setLevel(logging.INFO)

dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table(os.environ['MATCHLIST_TABLE'])


def insert(table, match_list):
    '''
    put items into specified DynamoDB table.
    '''
    with table.batch_writer() as batch:
        for match in match_list:
            logger.info('put match info into the table: {}'.format(match))
            batch.put_item({k: v for k, v in match.items()})


def is_json(json_str):
    '''
    judge if json_str is valid.
    '''
    result = False
    try:
        json.loads(json_str)
        result = True
    except json.JSONDecodeError as jde:
        logger.info('got invalid response json, retrying.')

    return result


def sleep():
    '''
    sleep for 1~10 secs (randomly)
    '''
    sec = random.randint(1, 10)
    time.sleep(sec)


def shorten(string):
    '''
    shorten strings for visibility in Google Calendar.
    '''
    shorten_string = string
    abbrs = {
        'VALORANT Champions Tour': 'VCT',
        'Last Chance Qualifier': 'LCQ',
        '2022 ': '',
        '2023 ': '',
        'North America': 'NA'
    }

    if shorten_string is None:
        shorten_string = ''
    else:
        for k, v in abbrs.items():
            shorten_string = shorten_string.replace(k, v)

    return shorten_string


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


def map_flag_to_region(flag):
    '''
    get flag(usually it shows country) indicator and return region
    '''

    region_map = {
        'au': 'APAC', 'ar': 'BR_LATAM', 'at': 'EMEA', 'ba': 'EMEA', 'be': 'EMEA', 'br': 'BR_LATAM',
        'ca': 'NA', 'ch': 'EMEA', 'cl': 'BR_LATAM', 'cn': 'EAST_ASIA', 'cz': 'EMEA', 'de': 'EMEA',
        'dk': 'EMEA', 'ee': 'EMEA', 'eg': 'EMEA', 'es': 'EMEA', 'fi': 'EMEA', 'fr': 'EMEA',
        'gb': 'EMEA', 'gr': 'EMEA', 'hk': 'APAC', 'hu': 'EMEA', 'hr': 'EMEA', 'id': 'APAC',
        'ie': 'EMEA', 'il': 'EMEA', 'in': 'APAC', 'is': 'EMEA', 'it': 'EMEA', 'jp': 'EAST_ASIA',
        'kh': 'APAC', 'kr': 'EAST_ASIA', 'kw': 'EMEA', 'lt': 'EMEA', 'ma': 'EMEA', 'me': 'EMEA', 'mk': 'EMEA', 'my': 'APAC',
        'no': 'EMEA', 'pe': 'BR_LATAM', 'ph': 'APAC', 'pl': 'EMEA', 'pt': 'EMEA', 'ro': 'EMEA',
        'rs': 'EMEA', 'sa': 'APAC', 'se': 'EMEA', 'sg': 'APAC', 'si': 'EMEA', 'th': 'APAC',
        'tr': 'EMEA', 'tw': 'APAC', 'ua': 'EMEA', 'us': 'NA', 'vn': 'APAC',
        'asia-pacific': 'APAC', 'benelux': 'EMEA', 'cis': 'EMEA', 'dach': 'EMEA',
        'east-asia': 'EAST_ASIA', 'eu': 'EMEA', 'latam': 'BR_LATAM', 'nordic': 'EMEA',
        'oce': 'APAC', 'south-asia': 'APAC', 'southeast-asia': 'APAC', 'usa-ca': 'NA', 'world': 'EMEA'
    }

    if flag in region_map:
        region = region_map[flag]
    else:
        logger.warning("no region map has found for the flag: {}".format(flag))
        region = ''

    return region


def get_matches():
    '''
    get match information from api endpoint
    '''

    endpoint = "https://api.thespike.gg/matches"
    days_to_get = 30

    logger.info('get matches list from: {}'.format(endpoint))

    headers = {"Content-Type": "application/json"}

    today = datetime.datetime.now()
    dates = [datetime.datetime.strftime(
        today + datetime.timedelta(days=d), '%Y-%m-%d') for d in range(days_to_get)]

    match_list = []

    for date in dates:
        logger.info('get matches list for a day: {}'.format(date))
        uri = endpoint + '?date=' + date

        # due to api endpoint's behabior, requests.get.json() sometimes fails
        # retry till valid json can be gotten
        upcoming_matches = ''
        while True:
            sleep()
            upcoming_matches = requests.get(uri, headers=headers)
            valid_json = is_json(upcoming_matches.text)

            if valid_json:
                break

        matches = upcoming_matches.json()

        # pick up information for each individual match
        for match in matches:

            # get region from flag indicator
            # skip if empty so that "Calendar Id" for Gcal cannot be determined
            flag = match['eventCountryFlag']
            region = map_flag_to_region(flag)

            if not region:
                continue

            match_id = str(match['id'])
            teams = [team['title'] for team in match['teams']]
            event_name = shorten(match['eventName'])

            # day x, upper/lower bracket, etc
            event_detail = shorten(match['matchName'])

            start_time = match['startTime']
            best_of = match['bestOf'] if type(match['bestOf']) is int else '3'
            # calculate when the match ends (and when deleted from DynamoDB)
            end_time, ttl = calc_match_end_time(start_time, int(best_of))

            # match url
            slugs = '-'.join([team['slug'] for team in match['teams']])
            prefix = "https://www.thespike.gg/match"
            match_uri = "/".join([prefix, slugs, match_id])

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
    match_list = get_matches()

    insert(table, match_list)

    return {
        'matches_count': len(match_list)
    }
