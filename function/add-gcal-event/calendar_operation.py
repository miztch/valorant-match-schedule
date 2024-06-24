import logging

from id_generators import generate_event_id
from region import (
    get_registered_regions,
    get_regions_to_register,
    map_region_to_calendar,
)
from dynamodb_operation import (
    deserialize,
    get_registered_calendars,
    get_gcal_event_id,
    put_event_id,
    delete_event_id,
)
from gcal_service import (
    get_gcal_credentials,
    insert_gcal_event,
    update_gcal_event,
    delete_gcal_event,
)

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def generate_event_body(item):
    """
    assemble json for creating / updating Google Calendar event
    """
    event_body = {
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

    return event_body


def populate_event(service, item, regions):
    """
    add or update calendar events
    """
    for region in regions:
        calendar_id = map_region_to_calendar(region)

        body = generate_event_body(item)
        event_id = get_gcal_event_id(item["match_id"], calendar_id)

        if event_id:
            logger.info("update existing event in calendar %s: %s", region, body)
            update_gcal_event(service, calendar_id, body, event_id)
        else:
            event_id = generate_event_id(item["match_id"])
            body["id"] = event_id
            logger.info("insert new event into calendar %s: %s", region, body)
            insert_gcal_event(service, calendar_id, body)

        put_event_id(item["match_id"], calendar_id, event_id, item["ttl"])


def delete_event(service, item, regions, registered_regions):
    """
    delete unnecessary calendar events
    """
    # if the event is registered at least in 1 calendar, pass to the next step
    # registered but not in regions given, remove it
    for region in registered_regions:
        if region not in regions:
            calendar_id = map_region_to_calendar(region)
            # Outbox DynamoDB Table has record = already registered
            event_id = get_gcal_event_id(item["match_id"], calendar_id)
            logger.info("delete existing event from calendar %s: %s", region, event_id)

            if event_id:
                delete_gcal_event(service, calendar_id, event_id)
                delete_event_id(item["match_id"], calendar_id)


def update_match_on_calendars(records):
    """
    update Google Calendar with adding/updating/removing the given match records
    """
    service = get_gcal_credentials()

    for record in records:
        # progress if not 'REMOVE' action
        # (do nothing when that and continue to the next record)
        if record["eventName"] == "REMOVE":
            logger.info("no action for REMOVE event")
            continue

        # DynamoDB JSON -> Python dict
        image = record["dynamodb"]["NewImage"]
        item = deserialize(image)

        # get registered regions and regions to register
        registered_regions = get_registered_regions(item)
        regions_to_register = get_regions_to_register(item)

        # update calendars
        delete_event(service, item, regions_to_register, registered_regions)
        populate_event(service, item, regions_to_register)
