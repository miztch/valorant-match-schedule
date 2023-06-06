countries = {
    'au': 'APAC', 'ar': 'BR_LATAM', 'at': 'EMEA', 'ba': 'EMEA', 'be': 'EMEA', 'br': 'BR_LATAM',
    'ca': 'NA', 'ch': 'EMEA', 'cl': 'BR_LATAM', 'cn': 'EAST_ASIA', 'cz': 'EMEA', 'de': 'EMEA',
    'dk': 'EMEA', 'ee': 'EMEA', 'eg': 'EMEA', 'es': 'EMEA', 'fi': 'EMEA', 'fr': 'EMEA', 'gb': 'EMEA',
    'gr': 'EMEA', 'hk': 'APAC', 'hu': 'EMEA', 'hr': 'EMEA', 'id': 'APAC', 'ie': 'EMEA', 'il': 'EMEA',
    'in': 'APAC', 'iq': 'EMEA', 'is': 'EMEA', 'it': 'EMEA', 'jp': 'EAST_ASIA', 'kh': 'APAC',
    'kr': 'EAST_ASIA', 'kw': 'EMEA', 'lt': 'EMEA', 'ma': 'EMEA', 'me': 'EMEA', 'mk': 'EMEA',
    'my': 'APAC', 'no': 'EMEA', 'pe': 'BR_LATAM', 'ph': 'APAC', 'pl': 'EMEA', 'pt': 'EMEA',
    'ro': 'EMEA', 'rs': 'EMEA', 'sa': 'EMEA', 'se': 'EMEA', 'sg': 'APAC', 'si': 'EMEA',
    'th': 'APAC', 'tr': 'EMEA', 'tw': 'APAC', 'ua': 'EMEA', 'us': 'NA', 'vn': 'APAC',
    'asia-pacific': 'APAC', 'benelux': 'EMEA', 'cis': 'EMEA', 'dach': 'EMEA', 'east-asia': 'EAST_ASIA', 'eu': 'EMEA',
    'latam': 'BR_LATAM', 'nordic': 'EMEA', 'oce': 'APAC', 'south-asia': 'APAC', 'southeast-asia': 'APAC', 'usa-ca': 'NA'
}

compellations = [
    {
        'APAC': [
            'SEA',
            'Asia Pacific',
            'Asia-Pacific',
            'Southeast Asia',
            'Oceania'
        ]
    },
    {
        'BR_LATAM': [
            'LATAM',
            'Latin America',
            'LA-N',
            'LA-S' 
        ]
    },
    {
        'EAST_ASIA': [
            'East Asia',
            'KR/JP'
        ]
    },
    {
        'EMEA': [
            'Europe',
            'Europian',
            'CIS',
            'MENA',
            'Middle East',
            'Arab',
            'Arabia',
            'Arabian',
            'Arabic',
            'Africa',
            'ESTAZ',
            'GLA'
        ]
    },
    {
        'NA': []
    }
]

# event name is after shortened
international_events = [
    'Champions Tour 2023: EMEA League',
    'Champions Tour 2023: Pacific League',
    'Champions Tour 2023: Americas League',
    'Champions Tour 2023: Masters Tokyo',
    'Champions Tour 2023: Pacific LCQ',
    'Champions Tour 2023: EMEA LCQ',
    'Champions Tour 2023: Americas LCQ',
    'Valorant Champions 2023'
]

abbrs = {
    'VALORANT Champions Tour': 'VCT',
    'Last Chance Qualifier': 'LCQ',
    'North America': 'NA'
}

headers = {"Content-Type": "application/json"}
