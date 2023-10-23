import json
import random
import time

import constants


def sleep():
    """
    sleep for 1~10 secs (randomly)
    """
    sec = random.randint(1, 10)
    time.sleep(sec)


def is_json(json_str):
    """
    judge if json_str is valid.
    """
    result = False
    try:
        json.loads(json_str)
        result = True
    except json.JSONDecodeError as jde:
        logger.info("got invalid response json, retrying.")

    return result


def shorten(string):
    """
    shorten strings for visibility in Google Calendar.
    """
    shorten_string = string
    abbrs = constants.abbrs

    if shorten_string is None:
        shorten_string = ""
    else:
        for k, v in abbrs.items():
            shorten_string = shorten_string.replace(k, v)

    return shorten_string
