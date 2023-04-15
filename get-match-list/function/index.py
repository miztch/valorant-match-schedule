import datetime
import json
import logging
import os

import boto3

logger = logging.getLogger()
logger.setLevel(logging.INFO)

sqs = boto3.client('sqs')
queue_url = os.environ['FETCH_DAILY_MATCHES_QUEUE_URL']


def publish_dates():
    '''
    publish dates to get match information to the queue
    '''

    days_to_get = int(os.environ['DAYS_TO_GET'])
    today = datetime.datetime.now()
    dates = [datetime.datetime.strftime(
        today + datetime.timedelta(days=d), '%Y-%m-%d') for d in range(days_to_get)]

    base_delay_seconds = int(os.environ['BASE_DELAY_SECONDS'])
    for i, date in zip(range(days_to_get), dates):
        logger.info('request to fetch match list for the day: {}'.format(date))

        payload = {'date': date}
        message = json.dumps(payload)

        response = sqs.send_message(
            QueueUrl=queue_url,
            MessageBody=message,
            DelaySeconds=i*base_delay_seconds
        )

        logger.info('message sent. queue: {} response: {}'.format(
            queue_url, response))


def lambda_handler(event, context):
    publish_dates()
