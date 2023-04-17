import json
import logging
import os
import random

import boto3
import google.auth
import googleapiclient.discovery
from boto3.dynamodb.conditions import Key
from boto3.dynamodb.types import TypeDeserializer

logger = logging.getLogger()
logger.setLevel(logging.INFO)

service_account_id = os.environ['GOOGLE_SERVICE_ACCOUNT_ID']

dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table(os.environ['OUTBOX_TABLE'])

deserializer = TypeDeserializer()


def get_gcal_credentials():
    scopes = ['https://www.googleapis.com/auth/calendar']

    credentials = google.auth.load_credentials_from_file(
        './google_key.json', scopes)[0]

    service = googleapiclient.discovery.build(
        'calendar', 'v3', credentials=credentials)

    return service


def map_region_to_calendar(region):
    '''
    returns calendar_id for the event
    '''

    calendars = {
        'APAC': os.environ['CALENDAR_ID_APAC'],
        'BR_LATAM': os.environ['CALENDAR_ID_BR_LATAM'],
        'EAST_ASIA': os.environ['CALENDAR_ID_EAST_ASIA'],
        'EMEA': os.environ['CALENDAR_ID_EMEA'],
        'NA': os.environ['CALENDAR_ID_NA']
    }

    calendar_id = calendars[region]
    return calendar_id


def deserialize(image):
    '''
    DynamoDB-formatted JSON -> Python dict
    '''
    d = {}
    for key in image:
        d[key] = deserializer.deserialize(image[key])
    return d


def put_event_id(match_id, calendar_id, event_id, matchlist_ttl):
    '''
    put item into outbox DynamoDB Table including
        - match_id
        - calendar id (of Google Calendar)
        - event_id (of Google Calendar)
        - ttl
    '''

    # outbox table ttl: matchlist_ttl + 30 days
    outbox_ttl = matchlist_ttl + 60 * 60 * 24 * 30

    table.put_item(
        Item={
            'match_id': match_id,
            'calendar_id': calendar_id,
            'event_id': event_id,
            'ttl': outbox_ttl
        }
    )


def if_gcal_event_registered(match_id):
    '''
    two functions
      - check if the Google Calendar event is already registered
        - outbox DynamoDB table has its status as an item
      - get event_id of Google Calendar from match_id in outbox DynamoDB Table
    '''
    record = table.get_item(Key={'match_id': match_id})
    if 'Item' in record:
        logger.info(
            'match id: {} found to already be registered'.format(match_id))
        return True, record['Item']['event_id']
    else:
        logger.info(
            'match id: {} was not found to be registered yet'.format(match_id))
        return False, ''


def assemble_gcal_event_json(action, item):
    '''
    assemble json for creating / updating Google Calendar event
    '''
    # event_id is needed only for "ADD" action
    # format: "match" + match_id + "0" + hash value
    # this length has no decent reason
    hash_length = 16 - len(str(item['match_id']))
    event_id = {
        'id': 'match{}0{}'.format(item['match_id'], str(random.randrange(10**hash_length, 10**(hash_length+1))))
    }
    detail = {
        'summary': '{} - {} | {} - {}'.format(item['team_home'], item['team_away'], item['event_name'], item['event_detail']),
        'description': item['match_uri'],
        'start': {
            'dateTime': item['start_time'],
            'timeZone': 'Etc/UTC',
        },
        'end': {
            'dateTime': item['end_time'],
            'timeZone': 'Etc/UTC',
        }
    }

    if action == "ADD":
        return dict(**event_id, **detail)
    else:
        return detail


def add_gcal_event(service_account_id, calendar_id, item):
    '''
    add new Google Calendar event
    '''
    service = get_gcal_credentials()

    body = assemble_gcal_event_json("ADD", item)

    try:
        logger.info('insert new event: {}'.format(body))
        result = service.events().insert(calendarId=calendar_id, body=body).execute()
        put_event_id(item['match_id'], calendar_id, body['id'], item['ttl'])
    except Exception as e:
        raise e

    return result


def update_gcal_event(service_account_id, calendar_id, item, event_id):
    '''
    update existing Google Calendar event
    '''
    service = get_gcal_credentials()

    body = assemble_gcal_event_json("UPDATE", item)

    try:
        logger.info('update existing event: {}'.format(body))
        result = service.events().update(calendarId=calendar_id,
                                         eventId=event_id, body=body).execute()
        put_event_id(item['match_id'], calendar_id, event_id, item['ttl'])
    except Exception as e:
        raise e

    return result


def lambda_handler(event, context):

    records = event['Records']
    logger.info(records)

    for record in records:
        # progress if not 'REMOVE' action
        # (do nothing when that)
        if record['eventName'] != 'REMOVE':

            # DynamoDB JSON -> Python dict
            image = record['dynamodb']['NewImage']
            item = deserialize(image)

            calendar_id = map_region_to_calendar(item['region'])

            try:
                # Outbox DynamoDB Table has record = already registered
                already_registered, event_id = if_gcal_event_registered(
                    item['match_id'])

                if already_registered:
                    update_gcal_event(service_account_id,
                                      calendar_id, item, event_id)
                else:
                    add_gcal_event(service_account_id, calendar_id, item)

            except Exception as e:
                raise e
