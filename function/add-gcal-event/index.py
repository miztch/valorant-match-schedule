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

dynamodb = boto3.resource("dynamodb")
table = dynamodb.Table(os.environ["OUTBOX_TABLE"])

deserializer = TypeDeserializer()


def get_gcal_credentials():
    scopes = ["https://www.googleapis.com/auth/calendar"]

    credentials = google.auth.load_credentials_from_file(
        "./service_account_key.json", scopes
    )[0]

    service = googleapiclient.discovery.build(
        "calendar", "v3", credentials=credentials, cache_discovery=False
    )

    return service


def map_region_to_calendar(region):
    """
    returns calendar_id for the event
    """

    calendar_id = os.environ[f"CALENDAR_ID_{region}"]
    return calendar_id


def map_calendar_to_region(calendar_id):
    """
    returns region for the calendar
    """

    for key, value in os.environ.items():
        if key.startswith("CALENDAR_ID_") and value == calendar_id:
            return key.replace("CALENDAR_ID_", "")

    return None


def deserialize(image):
    """
    DynamoDB-formatted JSON -> Python dict
    """
    d = {}
    for key in image:
        d[key] = deserializer.deserialize(image[key])
    return d


def put_event_id(match_id, calendar_id, event_id, matchlist_ttl):
    """
    put item into outbox DynamoDB Table including
        - match_id
        - calendar id (of Google Calendar)
        - event_id (of Google Calendar)
        - ttl
    """

    # outbox table ttl: matchlist_ttl + 30 days
    outbox_ttl = matchlist_ttl + 60 * 60 * 24 * 30

    table.put_item(
        Item={
            "match_id": match_id,
            "calendar_id": calendar_id,
            "event_id": event_id,
            "ttl": outbox_ttl,
        }
    )


def delete_event_id(match_id, calendar_id):
    """
    delete item from outbox DynamoDB Table
    """

    table.delete_item(Key={"match_id": match_id, "calendar_id": calendar_id})

    return None


def get_registered_calendars(match_id):
    """
    get calendars that the event is registered in
    """
    try:
        response = table.query(
            KeyConditionExpression=Key("match_id").eq(match_id),
        )
        calendars = [item["calendar_id"] for item in response["Items"]]
    except Exception as e:
        raise e
    else:
        return calendars


def get_gcal_event_id(match_id, calendar_id):
    """
    get event_id of Google Calendar from match_id in outbox DynamoDB Table
    """

    record = table.get_item(Key={"match_id": match_id, "calendar_id": calendar_id})

    if "Item" in record:
        logger.info(
            "match id: %s found in calendar id: %s",
            match_id,
            calendar_id,
        )
        return record["Item"]["event_id"]
    else:
        logger.info(
            "match id: %s not found in calendar id: %s yet",
            match_id,
            calendar_id,
        )
        return ""


def generate_event_id(match_id):
    """
    generate event_id for Google Calendar
    """
    # format: "match" + match_id + "0" + hash value
    # this length has no decent reason
    hash_length = 16 - len(str(match_id))
    random_string = str(random.randrange(10**hash_length, 10 ** (hash_length + 1)))

    event_id = f"match{match_id}0{random_string}"

    return event_id


def assemble_gcal_event_detail(item):
    """
    assemble json for creating / updating Google Calendar event
    """
    detail = {
        "summary": f"{item['team_home']} - {item['team_away']} | { item['event_name']} - {item['event_detail']}",
        "description": item["match_uri"],
        "start": {
            "dateTime": item["start_time"],
            "timeZone": "Etc/UTC",
        },
        "end": {
            "dateTime": item["end_time"],
            "timeZone": "Etc/UTC",
        },
    }

    return detail


def add_gcal_event(service, calendar_id, item):
    """
    add new Google Calendar event
    """
    body = assemble_gcal_event_detail(item)

    event_id = generate_event_id(item["match_id"])
    body["id"] = event_id

    try:
        logger.info("insert new event: %s", body)
        result = service.events().insert(calendarId=calendar_id, body=body).execute()
        put_event_id(item["match_id"], calendar_id, event_id, item["ttl"])
    except Exception as e:
        raise e

    return result


def update_gcal_event(service, calendar_id, item, event_id):
    """
    update existing Google Calendar event
    """
    body = assemble_gcal_event_detail(item)

    try:
        logger.info("update existing event: %s", body)
        result = (
            service.events()
            .update(calendarId=calendar_id, eventId=event_id, body=body)
            .execute()
        )
        put_event_id(item["match_id"], calendar_id, event_id, item["ttl"])
    except Exception as e:
        raise e

    return result


def delete_gcal_event(service, calendar_id, item, event_id):
    """
    delete existing Google Calendar event
    """

    try:
        logger.info("delete existing event: %s", event_id)
        result = (
            service.events().delete(calendarId=calendar_id, eventId=event_id).execute()
        )
        delete_event_id(item["match_id"], calendar_id)
    except Exception as e:
        raise e

    return None


def lambda_handler(event, context):
    records = event["Records"]
    logger.info(records)

    service = get_gcal_credentials()

    for record in records:
        # progress if not 'REMOVE' action
        # (do nothing when that)
        if record["eventName"] == "REMOVE":
            logger.info("no action for REMOVE event")
            return

        # DynamoDB JSON -> Python dict
        image = record["dynamodb"]["NewImage"]
        item = deserialize(image)

        # first, remove unnecessary calendar events. defined by "regions" for a match.
        # check which calendar the event is registered in. returns a list of calendars
        registered_calendars = get_registered_calendars(item["match_id"])
        # get registered regions from given calendar_ids
        registered_regions = [
            map_calendar_to_region(calendar)
            for calendar in registered_calendars
            if map_calendar_to_region(calendar) is not None
        ]
        logger.info(
            "match_id: %s is registered in calendar(s): %s",
            item["match_id"],
            registered_regions,
        )

        # item['region'] can be like "EMEA" or "EMEA#INTERNATIONAL"
        # if international event, add event to two calendars
        regions = item["region"].split("#")
        logger.info(
            "match_id: %s is for calendar(s): %s",
            item["match_id"],
            regions,
        )

        # if the event is registered at least in 1 calendar, pass to the next step
        # registered but not in given regions, remove it
        for region in registered_regions:
            if region not in regions:
                calendar_id = map_region_to_calendar(region)
                try:
                    # Outbox DynamoDB Table has record = already registered
                    event_id = get_gcal_event_id(item["match_id"], calendar_id)

                    if event_id:
                        delete_gcal_event(service, calendar_id, item, event_id)
                except Exception as e:
                    raise e

        # second, add or update calendar events.
        for region in regions:
            calendar_id = map_region_to_calendar(region)
            try:
                event_id = get_gcal_event_id(item["match_id"], calendar_id)

                if event_id:
                    update_gcal_event(service, calendar_id, item, event_id)
                else:
                    add_gcal_event(service, calendar_id, item)

            except Exception as e:
                raise e
