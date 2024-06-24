import logging

from calendar_operation import update_match_on_calendars

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    records = event["Records"]
    logger.info(records)

    update_match_on_calendars(records)
