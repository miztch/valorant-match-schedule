import logging
from boto3.dynamodb.conditions import Key
from boto3.dynamodb.types import TypeDeserializer
from dynamodb_service import (
    query,
    put_item,
    get_item,
    delete_item,
)

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def deserialize(image):
    """
    DynamoDB-formatted JSON -> Python dict
    """
    deserializer = TypeDeserializer()

    d = {}
    for key in image:
        d[key] = deserializer.deserialize(image[key])
    return d


def get_registered_calendars(match_id):
    """
    get calendars that the event is registered in
    """
    try:
        response = query(
            KeyConditionExpression=Key("match_id").eq(match_id),
        )
        calendars = [item["calendar_id"] for item in response["Items"]]
    except Exception as e:
        logger.error(e)
    else:
        return calendars


def get_gcal_event_id(match_id, calendar_id):
    """
    get event_id of Google Calendar from match_id in outbox DynamoDB Table
    """
    try:
        record = get_item(Key={"match_id": match_id, "calendar_id": calendar_id})
    except Exception as e:
        logger.error(e)
        return

    if "Item" in record:
        return record["Item"]["event_id"]
    else:
        return ""


def put_event_id(match_id, calendar_id, event_id, ttl):
    """
    put item into outbox DynamoDB Table including
        - match_id
        - calendar id (of Google Calendar)
        - event_id (of Google Calendar)
        - ttl
    """
    # outbox table ttl: matchlist_ttl + 30 days
    outbox_ttl = ttl + 60 * 60 * 24 * 30

    try:
        put_item(
            Item={
                "match_id": match_id,
                "calendar_id": calendar_id,
                "event_id": event_id,
                "ttl": outbox_ttl,
            }
        )
    except Exception as e:
        logger.error(e)


def delete_event_id(match_id, calendar_id):
    """
    delete item from outbox DynamoDB Table
    """
    try:
        delete_item(Key={"match_id": match_id, "calendar_id": calendar_id})
    except Exception as e:
        logger.error(e)
