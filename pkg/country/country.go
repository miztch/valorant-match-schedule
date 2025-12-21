// Package country provides country code to region mapping based on ISO 3166-1 alpha-2 and ccTLD standards.
// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
// Countries are assigned to one of four regions: EMEA, AMERICAS, PACIFIC, or CHINA.
// https://competitiveops.riotgames.com/en-US/VALORANT
package country

// CountryInfo represents information about a country including its name and region.
type CountryInfo struct {
	Name   string
	Region string
}

// Countries is a map where the key is the ISO 3166-1 alpha-2 country code
var Countries = map[string]CountryInfo{
	"ad": {Name: "Andorra", Region: "EMEA"},
	"ae": {Name: "United Arab Emirates", Region: "EMEA"},
	"af": {Name: "Afghanistan", Region: "EMEA"},
	"ag": {Name: "Antigua and Barbuda", Region: "AMERICAS"},
	"ai": {Name: "Anguilla", Region: "AMERICAS"},
	"al": {Name: "Albania", Region: "EMEA"},
	"am": {Name: "Armenia", Region: "EMEA"},
	"ao": {Name: "Angola", Region: "EMEA"},
	"aq": {Name: "Antarctica", Region: "PACIFIC"},
	"ar": {Name: "Argentina", Region: "AMERICAS"},
	"as": {Name: "American Samoa", Region: "PACIFIC"},
	"at": {Name: "Austria", Region: "EMEA"},
	"au": {Name: "Australia", Region: "PACIFIC"},
	"aw": {Name: "Aruba", Region: "AMERICAS"},
	"ax": {Name: "Åland Islands", Region: "EMEA"},
	"az": {Name: "Azerbaijan", Region: "EMEA"},
	"ba": {Name: "Bosnia and Herzegovina", Region: "EMEA"},
	"bb": {Name: "Barbados", Region: "AMERICAS"},
	"bd": {Name: "Bangladesh", Region: "PACIFIC"},
	"be": {Name: "Belgium", Region: "EMEA"},
	"bf": {Name: "Burkina Faso", Region: "EMEA"},
	"bg": {Name: "Bulgaria", Region: "EMEA"},
	"bh": {Name: "Bahrain", Region: "EMEA"},
	"bi": {Name: "Burundi", Region: "EMEA"},
	"bj": {Name: "Benin", Region: "EMEA"},
	"bl": {Name: "Saint Barthélemy", Region: "AMERICAS"},
	"bm": {Name: "Bermuda", Region: "AMERICAS"},
	"bn": {Name: "Brunei", Region: "PACIFIC"},
	"bo": {Name: "Bolivia", Region: "AMERICAS"},
	"bq": {Name: "Bonaire, Sint Eustatius and Saba", Region: "AMERICAS"},
	"br": {Name: "Brazil", Region: "AMERICAS"},
	"bs": {Name: "Bahamas", Region: "AMERICAS"},
	"bt": {Name: "Bhutan", Region: "PACIFIC"},
	"bv": {Name: "Bouvet Island", Region: "EMEA"},
	"bw": {Name: "Botswana", Region: "EMEA"},
	"by": {Name: "Belarus", Region: "EMEA"},
	"bz": {Name: "Belize", Region: "AMERICAS"},
	"ca": {Name: "Canada", Region: "AMERICAS"},
	"cc": {Name: "Cocos (Keeling) Islands", Region: "PACIFIC"},
	"cd": {Name: "Democratic Republic of the Congo", Region: "EMEA"},
	"cf": {Name: "Central African Republic", Region: "EMEA"},
	"cg": {Name: "Congo", Region: "EMEA"},
	"ch": {Name: "Switzerland", Region: "EMEA"},
	"ci": {Name: "Côte d'Ivoire", Region: "EMEA"},
	"ck": {Name: "Cook Islands", Region: "PACIFIC"},
	"cl": {Name: "Chile", Region: "AMERICAS"},
	"cm": {Name: "Cameroon", Region: "EMEA"},
	"cn": {Name: "China", Region: "CHINA"},
	"co": {Name: "Colombia", Region: "AMERICAS"},
	"cr": {Name: "Costa Rica", Region: "AMERICAS"},
	"cu": {Name: "Cuba", Region: "AMERICAS"},
	"cv": {Name: "Cape Verde", Region: "EMEA"},
	"cw": {Name: "Curaçao", Region: "AMERICAS"},
	"cx": {Name: "Christmas Island", Region: "PACIFIC"},
	"cy": {Name: "Cyprus", Region: "EMEA"},
	"cz": {Name: "Czech Republic", Region: "EMEA"},
	"de": {Name: "Germany", Region: "EMEA"},
	"dj": {Name: "Djibouti", Region: "EMEA"},
	"dk": {Name: "Denmark", Region: "EMEA"},
	"dm": {Name: "Dominica", Region: "AMERICAS"},
	"do": {Name: "Dominican Republic", Region: "AMERICAS"},
	"dz": {Name: "Algeria", Region: "EMEA"},
	"ec": {Name: "Ecuador", Region: "AMERICAS"},
	"ee": {Name: "Estonia", Region: "EMEA"},
	"eg": {Name: "Egypt", Region: "EMEA"},
	"eh": {Name: "Western Sahara", Region: "EMEA"},
	"er": {Name: "Eritrea", Region: "EMEA"},
	"es": {Name: "Spain", Region: "EMEA"},
	"et": {Name: "Ethiopia", Region: "EMEA"},
	"fi": {Name: "Finland", Region: "EMEA"},
	"fj": {Name: "Fiji", Region: "PACIFIC"},
	"fk": {Name: "Falkland Islands (Malvinas)", Region: "AMERICAS"},
	"fm": {Name: "Micronesia", Region: "PACIFIC"},
	"fo": {Name: "Faroe Islands", Region: "EMEA"},
	"fr": {Name: "France", Region: "EMEA"},
	"ga": {Name: "Gabon", Region: "EMEA"},
	"gb": {Name: "United Kingdom", Region: "EMEA"},
	"gd": {Name: "Grenada", Region: "AMERICAS"},
	"ge": {Name: "Georgia", Region: "EMEA"},
	"gf": {Name: "French Guiana", Region: "AMERICAS"},
	"gg": {Name: "Guernsey", Region: "EMEA"},
	"gh": {Name: "Ghana", Region: "EMEA"},
	"gi": {Name: "Gibraltar", Region: "EMEA"},
	"gl": {Name: "Greenland", Region: "EMEA"},
	"gm": {Name: "Gambia", Region: "EMEA"},
	"gn": {Name: "Guinea", Region: "EMEA"},
	"gp": {Name: "Guadeloupe", Region: "AMERICAS"},
	"gq": {Name: "Equatorial Guinea", Region: "EMEA"},
	"gr": {Name: "Greece", Region: "EMEA"},
	"gs": {Name: "South Georgia and the South Sandwich Islands", Region: "EMEA"},
	"gt": {Name: "Guatemala", Region: "AMERICAS"},
	"gu": {Name: "Guam", Region: "PACIFIC"},
	"gw": {Name: "Guinea-Bissau", Region: "EMEA"},
	"gy": {Name: "Guyana", Region: "AMERICAS"},
	"hk": {Name: "Hong Kong", Region: "PACIFIC"},
	"hm": {Name: "Heard Island and McDonald Islands", Region: "PACIFIC"},
	"hn": {Name: "Honduras", Region: "AMERICAS"},
	"hr": {Name: "Croatia", Region: "EMEA"},
	"ht": {Name: "Haiti", Region: "AMERICAS"},
	"hu": {Name: "Hungary", Region: "EMEA"},
	"id": {Name: "Indonesia", Region: "PACIFIC"},
	"ie": {Name: "Ireland", Region: "EMEA"},
	"il": {Name: "Israel", Region: "EMEA"},
	"im": {Name: "Isle of Man", Region: "EMEA"},
	"in": {Name: "India", Region: "PACIFIC"},
	"io": {Name: "British Indian Ocean Territory", Region: "EMEA"},
	"iq": {Name: "Iraq", Region: "EMEA"},
	"ir": {Name: "Iran", Region: "EMEA"},
	"is": {Name: "Iceland", Region: "EMEA"},
	"it": {Name: "Italy", Region: "EMEA"},
	"je": {Name: "Jersey", Region: "EMEA"},
	"jm": {Name: "Jamaica", Region: "AMERICAS"},
	"jo": {Name: "Jordan", Region: "EMEA"},
	"jp": {Name: "Japan", Region: "PACIFIC"},
	"ke": {Name: "Kenya", Region: "EMEA"},
	"kg": {Name: "Kyrgyzstan", Region: "EMEA"},
	"kh": {Name: "Cambodia", Region: "PACIFIC"},
	"ki": {Name: "Kiribati", Region: "PACIFIC"},
	"km": {Name: "Comoros", Region: "EMEA"},
	"kn": {Name: "Saint Kitts and Nevis", Region: "AMERICAS"},
	"kp": {Name: "North Korea", Region: "PACIFIC"},
	"kr": {Name: "South Korea", Region: "PACIFIC"},
	"kw": {Name: "Kuwait", Region: "EMEA"},
	"ky": {Name: "Cayman Islands", Region: "AMERICAS"},
	"kz": {Name: "Kazakhstan", Region: "EMEA"},
	"la": {Name: "Laos", Region: "PACIFIC"},
	"lb": {Name: "Lebanon", Region: "EMEA"},
	"lc": {Name: "Saint Lucia", Region: "AMERICAS"},
	"li": {Name: "Liechtenstein", Region: "EMEA"},
	"lk": {Name: "Sri Lanka", Region: "PACIFIC"},
	"lr": {Name: "Liberia", Region: "EMEA"},
	"ls": {Name: "Lesotho", Region: "EMEA"},
	"lt": {Name: "Lithuania", Region: "EMEA"},
	"lu": {Name: "Luxembourg", Region: "EMEA"},
	"lv": {Name: "Latvia", Region: "EMEA"},
	"ly": {Name: "Libya", Region: "EMEA"},
	"ma": {Name: "Morocco", Region: "EMEA"},
	"mc": {Name: "Monaco", Region: "EMEA"},
	"md": {Name: "Moldova", Region: "EMEA"},
	"me": {Name: "Montenegro", Region: "EMEA"},
	"mf": {Name: "Saint Martin (French part)", Region: "AMERICAS"},
	"mg": {Name: "Madagascar", Region: "EMEA"},
	"mh": {Name: "Marshall Islands", Region: "PACIFIC"},
	"mk": {Name: "North Macedonia", Region: "EMEA"},
	"ml": {Name: "Mali", Region: "EMEA"},
	"mm": {Name: "Myanmar", Region: "PACIFIC"},
	"mn": {Name: "Mongolia", Region: "PACIFIC"},
	"mo": {Name: "Macao", Region: "PACIFIC"},
	"mp": {Name: "Northern Mariana Islands", Region: "PACIFIC"},
	"mq": {Name: "Martinique", Region: "AMERICAS"},
	"mr": {Name: "Mauritania", Region: "EMEA"},
	"ms": {Name: "Montserrat", Region: "AMERICAS"},
	"mt": {Name: "Malta", Region: "EMEA"},
	"mu": {Name: "Mauritius", Region: "EMEA"},
	"mv": {Name: "Maldives", Region: "PACIFIC"},
	"mw": {Name: "Malawi", Region: "EMEA"},
	"mx": {Name: "Mexico", Region: "AMERICAS"},
	"my": {Name: "Malaysia", Region: "PACIFIC"},
	"mz": {Name: "Mozambique", Region: "EMEA"},
	"na": {Name: "Namibia", Region: "EMEA"},
	"nc": {Name: "New Caledonia", Region: "PACIFIC"},
	"ne": {Name: "Niger", Region: "EMEA"},
	"nf": {Name: "Norfolk Island", Region: "PACIFIC"},
	"ng": {Name: "Nigeria", Region: "EMEA"},
	"ni": {Name: "Nicaragua", Region: "AMERICAS"},
	"nl": {Name: "Netherlands", Region: "EMEA"},
	"no": {Name: "Norway", Region: "EMEA"},
	"np": {Name: "Nepal", Region: "PACIFIC"},
	"nr": {Name: "Nauru", Region: "PACIFIC"},
	"nu": {Name: "Niue", Region: "PACIFIC"},
	"nz": {Name: "New Zealand", Region: "PACIFIC"},
	"om": {Name: "Oman", Region: "EMEA"},
	"pa": {Name: "Panama", Region: "AMERICAS"},
	"pe": {Name: "Peru", Region: "AMERICAS"},
	"pf": {Name: "French Polynesia", Region: "PACIFIC"},
	"pg": {Name: "Papua New Guinea", Region: "PACIFIC"},
	"ph": {Name: "Philippines", Region: "PACIFIC"},
	"pk": {Name: "Pakistan", Region: "PACIFIC"},
	"pl": {Name: "Poland", Region: "EMEA"},
	"pm": {Name: "Saint Pierre and Miquelon", Region: "AMERICAS"},
	"pn": {Name: "Pitcairn", Region: "PACIFIC"},
	"pr": {Name: "Puerto Rico", Region: "AMERICAS"},
	"ps": {Name: "Palestine", Region: "EMEA"},
	"pt": {Name: "Portugal", Region: "EMEA"},
	"pw": {Name: "Palau", Region: "PACIFIC"},
	"py": {Name: "Paraguay", Region: "AMERICAS"},
	"qa": {Name: "Qatar", Region: "EMEA"},
	"re": {Name: "Réunion", Region: "EMEA"},
	"ro": {Name: "Romania", Region: "EMEA"},
	"rs": {Name: "Serbia", Region: "EMEA"},
	"ru": {Name: "Russia", Region: "EMEA"},
	"rw": {Name: "Rwanda", Region: "EMEA"},
	"sa": {Name: "Saudi Arabia", Region: "EMEA"},
	"sb": {Name: "Solomon Islands", Region: "PACIFIC"},
	"sc": {Name: "Seychelles", Region: "EMEA"},
	"sd": {Name: "Sudan", Region: "EMEA"},
	"se": {Name: "Sweden", Region: "EMEA"},
	"sg": {Name: "Singapore", Region: "PACIFIC"},
	"sh": {Name: "Saint Helena, Ascension and Tristan da Cunha", Region: "EMEA"},
	"si": {Name: "Slovenia", Region: "EMEA"},
	"sj": {Name: "Svalbard and Jan Mayen", Region: "EMEA"},
	"sk": {Name: "Slovakia", Region: "EMEA"},
	"sl": {Name: "Sierra Leone", Region: "EMEA"},
	"sm": {Name: "San Marino", Region: "EMEA"},
	"sn": {Name: "Senegal", Region: "EMEA"},
	"so": {Name: "Somalia", Region: "EMEA"},
	"sr": {Name: "Suriname", Region: "AMERICAS"},
	"ss": {Name: "South Sudan", Region: "EMEA"},
	"st": {Name: "São Tomé and Príncipe", Region: "EMEA"},
	"sv": {Name: "El Salvador", Region: "AMERICAS"},
	"sx": {Name: "Sint Maarten (Dutch part)", Region: "AMERICAS"},
	"sy": {Name: "Syrian Arab Republic", Region: "EMEA"},
	"sz": {Name: "Eswatini (Swaziland)", Region: "EMEA"},
	"tc": {Name: "Turks and Caicos Islands", Region: "AMERICAS"},
	"td": {Name: "Chad", Region: "EMEA"},
	"tf": {Name: "French Southern Territories", Region: "EMEA"},
	"tg": {Name: "Togo", Region: "EMEA"},
	"th": {Name: "Thailand", Region: "PACIFIC"},
	"tj": {Name: "Tajikistan", Region: "EMEA"},
	"tk": {Name: "Tokelau", Region: "PACIFIC"},
	"tl": {Name: "Timor-Leste", Region: "PACIFIC"},
	"tm": {Name: "Turkmenistan", Region: "EMEA"},
	"tn": {Name: "Tunisia", Region: "EMEA"},
	"to": {Name: "Tonga", Region: "PACIFIC"},
	"tr": {Name: "Turkey", Region: "EMEA"},
	"tt": {Name: "Trinidad and Tobago", Region: "AMERICAS"},
	"tv": {Name: "Tuvalu", Region: "PACIFIC"},
	"tw": {Name: "Taiwan", Region: "PACIFIC"},
	"tz": {Name: "Tanzania", Region: "EMEA"},
	"ua": {Name: "Ukraine", Region: "EMEA"},
	"ug": {Name: "Uganda", Region: "EMEA"},
	"um": {Name: "United States Minor Outlying Islands", Region: "PACIFIC"},
	"us": {Name: "United States", Region: "AMERICAS"},
	"uy": {Name: "Uruguay", Region: "AMERICAS"},
	"uz": {Name: "Uzbekistan", Region: "EMEA"},
	"va": {Name: "Vatican City", Region: "EMEA"},
	"vc": {Name: "Saint Vincent and the Grenadines", Region: "AMERICAS"},
	"ve": {Name: "Venezuela", Region: "AMERICAS"},
	"vg": {Name: "Virgin Islands (British)", Region: "AMERICAS"},
	"vi": {Name: "Virgin Islands (U.S.)", Region: "AMERICAS"},
	"vn": {Name: "Vietnam", Region: "PACIFIC"},
	"vu": {Name: "Vanuatu", Region: "PACIFIC"},
	"wf": {Name: "Wallis and Futuna", Region: "PACIFIC"},
	"ws": {Name: "Samoa", Region: "PACIFIC"},
	"xk": {Name: "Kosovo", Region: "EMEA"},
	"ye": {Name: "Yemen", Region: "EMEA"},
	"yt": {Name: "Mayotte", Region: "EMEA"},
	"za": {Name: "South Africa", Region: "EMEA"},
	"zm": {Name: "Zambia", Region: "EMEA"},
	"zw": {Name: "Zimbabwe", Region: "EMEA"},
}

// SubAreas maps regions to lists of alternative names and sub-regions that can be used to identify the region.
var SubAreas = map[string][]string{
	"PACIFIC":  {"Pacific", "East Asia", "KR/JP", "SEA", "Asia Pacific", "Asia-Pacific", "Southeast Asia", "South Asia", "Oceania", "MY/SG"},
	"AMERICAS": {"Americas", "LATAM", "Latin America", "LA-N", "LA-S"},
	"EMEA": {
		"EMEA", "Europe", "Europian", "CIS", "MENA", "Middle East", "Arab", "Arabia",
		"Arabian", "Arabic", "Africa", "NORTH//EAST",
	},
	"CHINA": {"China", "CN", "Chinese"},
}

// Organizers maps regions to lists of known tournament organizers in those regions.
var Organizers = map[string][]string{
	"PACIFIC":  {"Global Esports"},
	"AMERICAS": {},
	"EMEA":     {"EMG", "ESTAZ", "GLA"},
	"CHINA":    {},
}

// InternationalEvents contains a list of known international tournament ids
var InternationalEvents = []int{
	// not VCT
	1752, // Red Bull Home Ground #4
	2171, // Red Bull Home Ground #5
	2449, // Esports World Cup 2025

	// Year 2023
	1189, // Champions Tour 2023: Americas League
	1190, // Champions Tour 2023: EMEA League
	1191, // Champions Tour 2023: Pacific League
	1494, // Champions Tour 2023: Masters Tokyo
	1657, // Valorant Champions 2023
	1658, // Champions Tour 2023: Americas Last Chance Qualifier
	1659, // Champions Tour 2023: EMEA Last Chance Qualifier
	1660, // Champions Tour 2023: Pacific Last Chance Qualifier
	1750, // Game Changers 2023 Championship: São Paulo

	// Year 2024
	1921, // Champions Tour 2024: Masters Madrid
	1923, // Champions Tour 2024: Americas Kickoff
	1924, // Champions Tour 2024: Pacific Kickoff
	1925, // Champions Tour 2024: EMEA Kickoff
	1926, // Champions Tour 2024: China Kickoff
	1998, // Champions Tour 2024: EMEA Stage 1
	1999, // Champions Tour 2024: Masters Shanghai
	2002, // Champions Tour 2024: Pacific Stage 1
	2004, // Champions Tour 2024: Americas Stage 1
	2005, // Champions Tour 2024: Pacific Stage 2
	2006, // Champions Tour 2024: China Stage 1
	2094, // Champions Tour 2024: EMEA Stage 2
	2095, // Champions Tour 2024: Americas Stage 2
	2096, // Champions Tour 2024: China Stage 2
	2097, // Valorant Champions 2024
	2124, // Game Changers 2024 Championship: Berlin

	// Year 2025
	2274, // VCT 2025: Americas Kickoff || Champions Tour 2025: Americas Kickoff
	2275, // VCT 2025: China Kickoff || Champions Tour 2025: China Kickoff
	2276, // VCT 2025: EMEA Kickoff || Champions Tour 2025: EMEA Kickoff
	2277, // VCT 2025: Pacific Kickoff || Champions Tour 2025: Pacific Kickoff
	2281, // Valorant Masters Bangkok 2025 || Champions Tour 2025: Masters Bangkok
	2282, // Valorant Masters Toronto 2025 || Champions Tour 2025: Masters Toronto
	2283, // Valorant Champions 2025
	2347, // VCT 2025: Americas Stage 1 || Champions Tour 2025: Americas Stage 1
	2359, // VCT 2025: China Stage 1 || Champions Tour 2025: China Stage 1
	2379, // VCT 2025: Pacific Stage 1 || Champions Tour 2025: Pacific Stage 1
	2380, // VCT 2025: EMEA Stage 1 || Champions Tour 2025: EMEA Stage 1
	2498, // VCT 2025: EMEA Stage 2
	2499, // VCT 2025: China Stage 2
	2500, // VCT 2025: Pacific Stage 2
	2501, // VCT 2025: Americas Stage 2
	2596, // Game Changers 2025: Championship Seoul

	// Year 2026
	2682, // VCT 2026: Americas Kickoff
	2683, // VCT 2026: Pacific Kickoff
	2684, // VCT 2026: EMEA Kickoff
	2685, // VCT 2026: China Kickoff
	2760, // Valorant Masters Santiago 2026
	2765, // Valorant Masters London 2026
	2766, // Valorant Champions 2026
	2775, // VCT 2026: Pacific Stage 1
	2776, // VCT 2026: Pacific Stage 2


}
