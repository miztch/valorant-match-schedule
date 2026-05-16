package domain

// VlrEvent represents an event from vlr.gg
type VlrEvent struct {
	Id          int
	Name        string
	CountryFlag string
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
	2860, // VCT 2026: Americas Stage 1
	2863, // VCT 2026: EMEA Stage 1
	2864, // VCT 2026: China Stage 1
}
