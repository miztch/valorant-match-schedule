import os
import boto3

dynamodb = boto3.resource("dynamodb")
table = dynamodb.Table(os.environ["OUTBOX_TABLE"])


def query(**kwargs):
    """
    query outbox DynamoDB Table
    """
    try:
        response = table.query(**kwargs)
    except Exception as e:
        raise e

    return response


def put_item(**kwargs):
    """
    put item into outbox DynamoDB Table
    """
    try:
        table.put_item(**kwargs)
    except Exception as e:
        raise e

    return None


def get_item(**kwargs):
    """
    get item from outbox DynamoDB Table
    """
    try:
        record = table.get_item(**kwargs)
    except Exception as e:
        raise e

    return record


def delete_item(**kwargs):
    """
    delete item from outbox DynamoDB Table
    """
    try:
        table.delete_item(**kwargs)
    except Exception as e:
        raise e

    return None
