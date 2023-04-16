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


def load_json(file_name):
    '''
    get json object from file and return as dict
    '''
    with open(file_name) as file:
        dict = json.load(file)

    return dict


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


def map_flag_to_region(flag, region_map):
    '''
    get flag(usually it shows country) indicator and return region
    '''

    map = [country for country in region_map if country['country_code'] == flag]

    if map:
        region = map[0]['region']
    elif len(map) >= 2:
        logger.warning(
            "region map returned multiple records for the flag: {}".format(flag))
        region = ''
    else:
        logger.warning("no region map has found for the flag: {}".format(flag))
        region = ''

    return region


def fetch_daily_matches(date):
    '''
    fetch match information for the specified day from api endpoint
    '''
    # configure datasource
    region_map = load_json("./countries.json")
    endpoint = "https://api.thespike.gg/matches"

    logger.info('fetch matches list from: {}'.format(endpoint))

    # assemble request URI and header
    logger.info('fetch matches list for a day: {}'.format(date))
    uri = endpoint + '?date=' + date
    headers = {"Content-Type": "application/json"}

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
    match_list = []
    for match in matches:

        # get region from flag indicator
        # skip if empty so that "Calendar Id" for Gcal cannot be determined
        flag = match['eventCountryFlag']
        region = map_flag_to_region(flag, region_map)

        if not region:
            continue

        match_id = match['id']
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
        match_uri = "/".join([prefix, slugs, str(match_id)])

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

    insert(table, match_list)

    return {
        'matches_count': len(match_list)
    }
