import random


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
