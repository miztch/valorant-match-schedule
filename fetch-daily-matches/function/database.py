import logging
import os

import boto3

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def _get_table():
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(os.environ['MATCHLIST_TABLE'])

    return table


def insert(match_list):
    '''
    put items into specified DynamoDB table.
    '''
    table = _get_table()

    with table.batch_writer() as batch:
        for match in match_list:
            logger.info('put match info into the table: {}'.format(match))
            batch.put_item({k: v for k, v in match.items()})
