import logging
import os

from dynamodb_operation import get_registered_calendars

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def get_registered_regions(item):
    """
    get region names of calendars that the event is registered in
    """
    registered_calendars = get_registered_calendars(item["match_id"])
    registered_regions = [
        map_calendar_to_region(calendar)
        for calendar in registered_calendars
        if map_calendar_to_region(calendar) is not None
    ]
    logger.info(
        "match_id: %s is in calendar(s): %s",
        item["match_id"],
        registered_regions,
    )

    return registered_regions


def get_regions_to_register(item):
    """
    get region names to register the event
    """
    # item['region'] can be like "EMEA" or "EMEA#INTERNATIONAL"
    # if international event, add event to two calendars
    regions = item["region"].split("#")
    logger.info(
        "match_id: %s is for calendar(s): %s",
        item["match_id"],
        regions,
    )

    return regions


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
