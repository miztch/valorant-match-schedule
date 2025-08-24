package application

import (
	"log/slog"
	"strings"
	"time"

	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
	"github.com/miztch/valorant-match-schedule/pkg/country"
)

const (
	vlrBaseURL    = "https://vlr.gg"
	offsetTTL     = 12 // Offset in hours for match TTL calculation
	bestOfDefault = 3  // Default bestOf value if not specified
)

func buildMatchURI(pagePath string) string {
	if pagePath == "" {
		return vlrBaseURL
	}
	if !strings.HasPrefix(pagePath, "/") {
		pagePath = "/" + pagePath
	}
	return vlrBaseURL + pagePath
}

// calcEndTime calculates the end time of a match based on its start time and bestOf value.
func calcEndTime(startTime string, bestOf int) (endTime string) {
	st, _ := time.Parse("2006-01-02T15:04:05-0700", startTime)
	var et time.Time
	if bestOf > 0 {
		et = st.Add(time.Duration(bestOf) * time.Hour)
	} else {
		et = st.Add(time.Duration(bestOfDefault) * time.Hour)
	}

	return et.Format("2006-01-02T15:04:05-0700")
}

// calcTTL calculates the TTL (Time To Live) for a match based on its end time.
func calcTTL(endTime string, offset int64) (ttl int64) {
	et, _ := time.Parse("2006-01-02T15:04:05-0700", endTime)
	// add offset in hours to end time
	return et.Add(time.Duration(offset) * time.Hour).Unix()
}

// Region
func mapFlagToRegion(countryFlag string) string {
	if info, exists := country.Countries[countryFlag]; exists {
		return info.Region
	}
	return ""
}

// containsAny checks if the eventName contains any of the keywords in the list.
func containsAny(eventName string, keywords []string) bool {
	for _, kw := range keywords {
		if strings.Contains(eventName, kw) {
			return true
		}
	}
	return false
}

func EstimateRegionByEvent(eventName string) string {
	for region, keywords := range country.SubAreas {
		if containsAny(eventName, keywords) {
			return region
		}
	}
	for region, orgs := range country.Organizers {
		if containsAny(eventName, orgs) {
			return region
		}
	}
	return ""
}

func isInternationalEvent(eventName string) bool {
	return containsAny(eventName, country.InternationalEvents)
}

// MatchToDTO converts domain.Match to dto.Match
func MatchToDTO(m domain.Match) dto.Match {
	endTime := calcEndTime(m.StartTime, m.BestOf)
	var region string
	c := m.EventCountryFlag
	if c == "" || c == "un" {
		region = EstimateRegionByEvent(m.EventName)
	} else {
		region = mapFlagToRegion(c)
	}

	if region == "" {
		slog.Warn("Region is empty for match", "matchID", m.Id, "eventName", m.EventName)
	}

	if isInternationalEvent(m.EventName) {
		region += "#INTERNATIONAL"
	}

	return dto.Match{
		MatchID:     m.Id,
		BestOf:      m.BestOf,
		EndTime:     endTime,
		EventDetail: m.Name,
		EventName:   m.EventName,
		MatchURI:    buildMatchURI(m.PagePath),
		Region:      region,
		StartTime:   m.StartTime,
		TeamAway:    m.Teams[1].Name,
		TeamHome:    m.Teams[0].Name,
		TTL:         calcTTL(endTime, offsetTTL),
	}
}
