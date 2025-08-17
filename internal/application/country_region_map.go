package application

// CountryToRegionMap is a mapping from country code to region string
// This mapping is based on the ISO 3166-1 standard (alpha-2) and the competitive regions for VALORANT.
// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
// https://competitiveops.riotgames.com/en-US/VALORANT
var CountryToRegionMap = map[string]string{
	"ad":             "EMEA",     // Andorra
	"ae":             "EMEA",     // United Arab Emirates
	"af":             "EMEA",     // Afghanistan
	"ag":             "AMERICAS", // Antigua and Barbuda
	"ai":             "AMERICAS", // Anguilla
	"al":             "EMEA",     // Albania
	"am":             "EMEA",     // Armenia
	"ao":             "EMEA",     // Angola
	"ar":             "AMERICAS", // Argentina
	"as":             "PACIFIC",  // American Samoa
	"at":             "EMEA",     // Austria
	"au":             "PACIFIC",  // Australia
	"aw":             "AMERICAS", // Aruba
	"ax":             "EMEA",     // Åland Islands
	"az":             "EMEA",     // Azerbaijan
	"ba":             "EMEA",     // Bosnia and Herzegovina
	"bb":             "AMERICAS", // Barbados
	"bd":             "PACIFIC",  // Bangladesh
	"be":             "EMEA",     // Belgium
	"bf":             "EMEA",     // Burkina Faso
	"bg":             "EMEA",     // Bulgaria
	"bh":             "EMEA",     // Bahrain
	"bi":             "EMEA",     // Burundi
	"bj":             "EMEA",     // Benin
	"bl":             "AMERICAS", // Saint Barthélemy
	"bm":             "AMERICAS", // Bermuda
	"bn":             "PACIFIC",  // Brunei
	"bo":             "AMERICAS", // Bolivia
	"bq":             "AMERICAS", // Bonaire, Sint Eustatius and Saba
	"br":             "AMERICAS", // Brazil
	"bs":             "AMERICAS", // Bahamas
	"bt":             "PACIFIC",  // Bhutan
	"bw":             "EMEA",     // Botswana
	"by":             "EMEA",     // Belarus
	"bz":             "AMERICAS", // Belize
	"ca":             "AMERICAS", // Canada
	"cc":             "PACIFIC",  // Cocos (Keeling) Islands
	"cd":             "EMEA",     // Democratic Republic of the Congo
	"cf":             "EMEA",     // Central African Republic
	"cg":             "EMEA",     // Congo
	"ch":             "EMEA",     // Switzerland
	"ci":             "EMEA",     // Côte d'Ivoire
	"ck":             "PACIFIC",  // Cook Islands
	"cl":             "AMERICAS", // Chile
	"cm":             "EMEA",     // Cameroon
	"cn":             "CHINA",    // China
	"co":             "AMERICAS", // Colombia
	"cr":             "AMERICAS", // Costa Rica
	"cu":             "AMERICAS", // Cuba
	"cv":             "EMEA",     // Cape Verde
	"cw":             "AMERICAS", // Curaçao
	"cx":             "PACIFIC",  // Christmas Island
	"cy":             "EMEA",     // Cyprus
	"cz":             "EMEA",     // Czech Republic
	"de":             "EMEA",     // Germany
	"dj":             "EMEA",     // Djibouti
	"dk":             "EMEA",     // Denmark
	"dm":             "AMERICAS", // Dominica
	"do":             "AMERICAS", // Dominican Republic
	"dz":             "EMEA",     // Algeria
	"ec":             "AMERICAS", // Ecuador
	"ee":             "EMEA",     // Estonia
	"eg":             "EMEA",     // Egypt
	"eh":             "EMEA",     // Western Sahara
	"er":             "EMEA",     // Eritrea
	"es":             "EMEA",     // Spain
	"et":             "EMEA",     // Ethiopia
	"fi":             "EMEA",     // Finland
	"fj":             "PACIFIC",  // Fiji
	"fk":             "AMERICAS", // Falkland Islands (Malvinas)
	"fm":             "PACIFIC",  // Micronesia
	"fo":             "EMEA",     // Faroe Islands
	"fr":             "EMEA",     // France
	"ga":             "EMEA",     // Gabon
	"gb":             "EMEA",     // United Kingdom
	"gd":             "AMERICAS", // Grenada
	"ge":             "EMEA",     // Georgia
	"gf":             "AMERICAS", // French Guiana
	"gg":             "EMEA",     // Guernsey
	"gh":             "EMEA",     // Ghana
	"gi":             "EMEA",     // Gibraltar
	"gl":             "EMEA",     // Greenland
	"gm":             "EMEA",     // Gambia
	"gn":             "EMEA",     // Guinea
	"gp":             "AMERICAS", // Guadeloupe
	"gq":             "EMEA",     // Equatorial Guinea
	"gr":             "EMEA",     // Greece
	"gt":             "AMERICAS", // Guatemala
	"gu":             "PACIFIC",  // Guam
	"gw":             "EMEA",     // Guinea-Bissau
	"gy":             "AMERICAS", // Guyana
	"hk":             "PACIFIC",  // Hong Kong
	"hn":             "AMERICAS", // Honduras
	"hr":             "EMEA",     // Croatia
	"ht":             "AMERICAS", // Haiti
	"hu":             "EMEA",     // Hungary
	"id":             "PACIFIC",  // Indonesia
	"ie":             "EMEA",     // Ireland
	"il":             "EMEA",     // Israel
	"im":             "EMEA",     // Isle of Man
	"in":             "PACIFIC",  // India
	"iq":             "EMEA",     // Iraq
	"ir":             "EMEA",     // Iran
	"is":             "EMEA",     // Iceland
	"it":             "EMEA",     // Italy
	"je":             "EMEA",     // Jersey
	"jm":             "AMERICAS", // Jamaica
	"jo":             "EMEA",     // Jordan
	"jp":             "PACIFIC",  // Japan
	"ke":             "EMEA",     // Kenya
	"kg":             "EMEA",     // Kyrgyzstan
	"kh":             "PACIFIC",  // Cambodia
	"ki":             "PACIFIC",  // Kiribati
	"km":             "EMEA",     // Comoros
	"kn":             "AMERICAS", // Saint Kitts and Nevis
	"kp":             "PACIFIC",  // North Korea
	"kr":             "PACIFIC",  // South Korea
	"kw":             "EMEA",     // Kuwait
	"ky":             "AMERICAS", // Cayman Islands
	"kz":             "EMEA",     // Kazakhstan
	"la":             "PACIFIC",  // Laos
	"lb":             "EMEA",     // Lebanon
	"lc":             "AMERICAS", // Saint Lucia
	"li":             "EMEA",     // Liechtenstein
	"lk":             "PACIFIC",  // Sri Lanka
	"lr":             "EMEA",     // Liberia
	"ls":             "EMEA",     // Lesotho
	"lt":             "EMEA",     // Lithuania
	"lu":             "EMEA",     // Luxembourg
	"lv":             "EMEA",     // Latvia
	"ly":             "EMEA",     // Libya
	"ma":             "EMEA",     // Morocco
	"mc":             "EMEA",     // Monaco
	"md":             "EMEA",     // Moldova
	"me":             "EMEA",     // Montenegro
	"mf":             "AMERICAS", // Saint Martin (French part)
	"mg":             "EMEA",     // Madagascar
	"mh":             "PACIFIC",  // Marshall Islands
	"mk":             "EMEA",     // North Macedonia
	"ml":             "EMEA",     // Mali
	"mm":             "PACIFIC",  // Myanmar
	"mn":             "PACIFIC",  // Mongolia
	"mo":             "PACIFIC",  // Macao
	"mp":             "PACIFIC",  // Northern Mariana Islands
	"mq":             "AMERICAS", // Martinique
	"mr":             "EMEA",     // Mauritania
	"ms":             "AMERICAS", // Montserrat
	"mt":             "EMEA",     // Malta
	"mu":             "EMEA",     // Mauritius
	"mv":             "PACIFIC",  // Maldives
	"mw":             "EMEA",     // Malawi
	"mx":             "AMERICAS", // Mexico
	"my":             "PACIFIC",  // Malaysia
	"mz":             "EMEA",     // Mozambique
	"na":             "EMEA",     // Namibia
	"nc":             "PACIFIC",  // New Caledonia
	"ne":             "EMEA",     // Niger
	"nf":             "PACIFIC",  // Norfolk Island
	"ng":             "EMEA",     // Nigeria
	"ni":             "AMERICAS", // Nicaragua
	"nl":             "EMEA",     // Netherlands
	"no":             "EMEA",     // Norway
	"np":             "PACIFIC",  // Nepal
	"nr":             "PACIFIC",  // Nauru
	"nu":             "PACIFIC",  // Niue
	"nz":             "PACIFIC",  // New Zealand
	"om":             "EMEA",     // Oman
	"pa":             "AMERICAS", // Panama
	"pe":             "AMERICAS", // Peru
	"pf":             "PACIFIC",  // French Polynesia
	"pg":             "PACIFIC",  // Papua New Guinea
	"ph":             "PACIFIC",  // Philippines
	"pk":             "PACIFIC",  // Pakistan
	"pl":             "EMEA",     // Poland
	"pm":             "AMERICAS", // Saint Pierre and Miquelon
	"pn":             "PACIFIC",  // Pitcairn
	"pr":             "AMERICAS", // Puerto Rico
	"ps":             "EMEA",     // Palestine
	"pt":             "EMEA",     // Portugal
	"pw":             "PACIFIC",  // Palau
	"py":             "AMERICAS", // Paraguay
	"qa":             "EMEA",     // Qatar
	"re":             "EMEA",     // Réunion
	"ro":             "EMEA",     // Romania
	"rs":             "EMEA",     // Serbia
	"ru":             "EMEA",     // Russia
	"rw":             "EMEA",     // Rwanda
	"sa":             "EMEA",     // Saudi Arabia
	"sb":             "PACIFIC",  // Solomon Islands
	"sc":             "EMEA",     // Seychelles
	"sd":             "EMEA",     // Sudan
	"se":             "EMEA",     // Sweden
	"sg":             "PACIFIC",  // Singapore
	"sh":             "EMEA",     // Saint Helena, Ascension and Tristan da Cunha
	"si":             "EMEA",     // Slovenia
	"sj":             "EMEA",     // Svalbard and Jan Mayen
	"sk":             "EMEA",     // Slovakia
	"sl":             "EMEA",     // Sierra Leone
	"sm":             "EMEA",     // San Marino
	"sn":             "EMEA",     // Senegal
	"so":             "EMEA",     // Somalia
	"sr":             "AMERICAS", // Suriname
	"ss":             "EMEA",     // South Sudan
	"st":             "EMEA",     // São Tomé and Príncipe
	"sv":             "AMERICAS", // El Salvador
	"sx":             "AMERICAS", // Sint Maarten (Dutch part)
	"sy":             "EMEA",     // Syrian Arab Republic
	"sz":             "EMEA",     // Eswatini (Swaziland)
	"tc":             "AMERICAS", // Turks and Caicos Islands
	"td":             "EMEA",     // Chad
	"tg":             "EMEA",     // Togo
	"th":             "PACIFIC",  // Thailand
	"tj":             "EMEA",     // Tajikistan
	"tk":             "PACIFIC",  // Tokelau
	"tl":             "PACIFIC",  // Timor-Leste
	"tm":             "EMEA",     // Turkmenistan
	"tn":             "EMEA",     // Tunisia
	"to":             "PACIFIC",  // Tonga
	"tr":             "EMEA",     // Turkey
	"tt":             "AMERICAS", // Trinidad and Tobago
	"tv":             "PACIFIC",  // Tuvalu
	"tw":             "PACIFIC",  // Taiwan
	"tz":             "EMEA",     // Tanzania
	"ua":             "EMEA",     // Ukraine
	"ug":             "EMEA",     // Uganda
	"us":             "AMERICAS", // United States
	"uy":             "AMERICAS", // Uruguay
	"uz":             "EMEA",     // Uzbekistan
	"va":             "EMEA",     // Vatican City
	"vc":             "AMERICAS", // Saint Vincent and the Grenadines
	"ve":             "AMERICAS", // Venezuela
	"vg":             "AMERICAS", // Virgin Islands (British)
	"vi":             "AMERICAS", // Virgin Islands (U.S.)
	"vn":             "PACIFIC",  // Vietnam
	"vu":             "PACIFIC",  // Vanuatu
	"wf":             "PACIFIC",  // Wallis and Futuna
	"ws":             "PACIFIC",  // Samoa
	"xk":             "EMEA",     // Kosovo
	"ye":             "EMEA",     // Yemen
	"yt":             "EMEA",     // Mayotte
	"za":             "EMEA",     // South Africa
	"zm":             "EMEA",     // Zambia
	"zw":             "EMEA",     // Zimbabwe
	"asia-pacific":   "PACIFIC",
	"benelux":        "EMEA",
	"cis":            "EMEA",
	"dach":           "EMEA",
	"east-asia":      "PACIFIC",
	"eu":             "EMEA",
	"latam":          "AMERICAS",
	"nordic":         "EMEA",
	"oce":            "PACIFIC",
	"south-asia":     "PACIFIC",
	"southeast-asia": "PACIFIC",
	"usa-ca":         "AMERICAS",
	// "aq":             "PACIFIC",  // Antarctica
	// "bv":             "EMEA",     // Bouvet Island
	// "gs":             "EMEA",     // South Georgia and the South Sandwich Islands
	// "hm":             "PACIFIC",  // Heard Island and McDonald Islands
	// "io":             "EMEA",     // British Indian Ocean Territory
	// "tf":             "EMEA",     // French Southern Territories
	// "um":             "PACIFIC",  // United States Minor Outlying Islands
}

var SubAreas = map[string][]string{
	"PACIFIC":  {"Pacific", "East Asia", "KR/JP", "SEA", "Asia Pacific", "Asia-Pacific", "Southeast Asia", "South Asia", "Oceania", "MY/SG"},
	"AMERICAS": {"Americas", "LATAM", "Latin America", "LA-N", "LA-S"},
	"EMEA": {
		"EMEA", "Europe", "Europian", "CIS", "MENA", "Middle East", "Arab", "Arabia",
		"Arabian", "Arabic", "Africa", "NORTH//EAST",
	},
	"CHINA": {"China", "CN", "Chinese"},
}

var Organizers = map[string][]string{
	"PACIFIC":  {"Global Esports"},
	"AMERICAS": {},
	"EMEA":     {"EMG", "ESTAZ", "GLA"},
	"CHINA":    {},
}

var InternationalEvents = []string{
	"Champions Tour 2023: EMEA League",
	"Champions Tour 2023: Pacific League",
	"Champions Tour 2023: Americas League",
	"Champions Tour 2023: Masters Tokyo",
	"Champions Tour 2023: Pacific Last Chance Qualifier",
	"Champions Tour 2023: EMEA Last Chance Qualifier",
	"Champions Tour 2023: Americas Last Chance Qualifier",
	"Valorant Champions 2023",
	"Game Changers 2023 Championship: São Paulo",
	"Red Bull Home Ground #4",
	"Champions Tour 2024: Americas Kickoff",
	"Champions Tour 2024: Pacific Kickoff",
	"Champions Tour 2024: EMEA Kickoff",
	"Champions Tour 2024: China Kickoff",
	"Champions Tour 2024: Masters Madrid",
	"Champions Tour 2024: Americas Stage 1",
	"Champions Tour 2024: EMEA Stage 1",
	"Champions Tour 2024: Pacific Stage 1",
	"Champions Tour 2024: China Stage 1",
	"Champions Tour 2024: Masters Shanghai",
	"Champions Tour 2024: Americas Stage 2",
	"Champions Tour 2024: EMEA Stage 2",
	"Champions Tour 2024: Pacific Stage 2",
	"Champions Tour 2024: China Stage 2",
	"Valorant Champions 2024",
	"Game Changers 2024 Championship: Berlin",
	"Red Bull Home Ground #5",
	"Champions Tour 2025: Americas Kickoff",
	"Champions Tour 2025: Pacific Kickoff",
	"Champions Tour 2025: EMEA Kickoff",
	"Champions Tour 2025: China Kickoff",
	"VCT 2025: Americas Kickoff",
	"VCT 2025: Pacific Kickoff",
	"VCT 2025: EMEA Kickoff",
	"VCT 2025: China Kickoff",
	"Champions Tour 2025: Masters Bangkok",
	"Valorant Masters Bangkok 2025",
	"Champions Tour 2025: Masters Toronto",
	"Valorant Masters Toronto 2025",
	"Champions Tour 2025: China Stage 1",
	"Champions Tour 2025: Americas Stage 1",
	"Champions Tour 2025: EMEA Stage 1",
	"Champions Tour 2025: Pacific Stage 1",
	"VCT 2025: China Stage 1",
	"VCT 2025: Americas Stage 1",
	"VCT 2025: EMEA Stage 1",
	"VCT 2025: Pacific Stage 1",
	"VCT 2025: China Stage 2",
	"VCT 2025: Americas Stage 2",
	"VCT 2025: EMEA Stage 2",
	"VCT 2025: Pacific Stage 2",
	"Valorant Champions 2025",
	"Esports World Cup 2025",
}
