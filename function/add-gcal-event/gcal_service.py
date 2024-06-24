import logging
import os

import google.auth
import googleapiclient.discovery


logger = logging.getLogger()
logger.setLevel(logging.INFO)


def get_gcal_credentials():
    """
    get Google Calendar credentials
    """
    scopes = ["https://www.googleapis.com/auth/calendar"]

    credentials = google.auth.load_credentials_from_file(
        "./service_account_key.json", scopes
    )[0]

    service = googleapiclient.discovery.build(
        "calendar", "v3", credentials=credentials, cache_discovery=False
    )

    return service


def insert_gcal_event(service, calendar_id, body):
    """
    add new Google Calendar event
    """
    try:
        result = service.events().insert(calendarId=calendar_id, body=body).execute()
    except Exception as e:
        raise e

    return result


def update_gcal_event(service, calendar_id, body, event_id):
    """
    update existing Google Calendar event
    """
    try:
        result = (
            service.events()
            .update(calendarId=calendar_id, eventId=event_id, body=body)
            .execute()
        )
    except Exception as e:
        raise e

    return result


def delete_gcal_event(service, calendar_id, event_id):
    """
    delete existing Google Calendar event
    """
    try:
        result = (
            service.events().delete(calendarId=calendar_id, eventId=event_id).execute()
        )
    except Exception as e:
        raise e

    return None
