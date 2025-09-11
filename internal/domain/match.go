package domain

import (
	"slices"

	"github.com/miztch/valorant-match-schedule/pkg/country"
)

// VlrMatch represents a match from vlr.gg
type VlrMatch struct {
	Id            int
	Name          string
	StartDate     string
	StartTime     string
	BestOf        int
	Teams         []Team
	PagePath      string
	EventPagePath string
}

// Match represents a match with associated event details
type Match struct {
	Id               int
	Name             string
	StartDate        string
	StartTime        string
	BestOf           int
	Teams            []Team
	PagePath         string
	EventId          int
	EventName        string
	EventCountryFlag string
}

// Team represents a team
type Team struct {
	Name string
}

// NewMatch creates a new match
func NewMatch(m VlrMatch, e VlrEvent) Match {
	return Match{
		Id:               m.Id,
		Name:             m.Name,
		StartDate:        m.StartDate,
		StartTime:        m.StartTime,
		BestOf:           m.BestOf,
		Teams:            m.Teams,
		PagePath:         m.PagePath,
		EventId:          e.Id,
		EventName:        e.Name,
		EventCountryFlag: e.CountryFlag,
	}
}

// isInternational checks if the event which provided match belongs to is an international event
func (m Match) IsInternational() bool {
	return slices.Contains(country.InternationalEvents, m.EventId)
}
