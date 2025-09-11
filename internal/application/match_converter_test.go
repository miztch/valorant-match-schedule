package application

import (
	"testing"

	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

func TestMatchToDTO(t *testing.T) {
	tests := []struct {
		name  string
		match domain.Match
		want  dto.Match
	}{
		{
			name: "Regional event",
			match: domain.Match{
				Id:               531976,
				Name:             "Main Event: Lower Round 1",
				StartDate:        "2025-08-18",
				StartTime:        "2025-08-18T21:00:00+0000",
				BestOf:           3,
				Teams:            []domain.Team{{Name: "Twisted Minds"}, {Name: "Team RA'AD"}},
				PagePath:         "/531976/twisted-minds-vs-team-raad-challengers-2025-mena-resilience-lan-finals-lr1",
				EventId:          "2577",
				EventName:        "Challengers 2025: MENA Resilience LAN Finals",
				EventCountryFlag: "sa",
			},
			want: dto.Match{
				MatchID:     531976,
				BestOf:      3,
				EndTime:     "2025-08-19T00:00:00+0000",
				EventDetail: "Main Event: Lower Round 1",
				EventName:   "Challengers 2025: MENA Resilience LAN Finals",
				MatchURI:    "https://vlr.gg/531976/twisted-minds-vs-team-raad-challengers-2025-mena-resilience-lan-finals-lr1",
				Region:      "EMEA",
				StartTime:   "2025-08-18T21:00:00+0000",
				TeamAway:    "Team RA'AD",
				TeamHome:    "Twisted Minds",
				TTL:         1755604800,
			},
		},
		{
			name: "International event",
			match: domain.Match{
				Id:               378829,
				Name:             "Playoffs: Grand Final",
				StartDate:        "2024-08-25",
				StartTime:        "2024-08-25T07:00:00+0000",
				BestOf:           5,
				Teams:            []domain.Team{{Name: "EDward Gaming"}, {Name: "Team Heretics"}},
				PagePath:         "/378829/edward-gaming-vs-team-heretics-valorant-champions-2024-gf",
				EventId:          "2097",
				EventName:        "Valorant Champions 2024",
				EventCountryFlag: "kr",
			},
			want: dto.Match{
				MatchID:     378829,
				BestOf:      5,
				EndTime:     "2024-08-25T12:00:00+0000",
				EventDetail: "Playoffs: Grand Final",
				EventName:   "Valorant Champions 2024",
				MatchURI:    "https://vlr.gg/378829/edward-gaming-vs-team-heretics-valorant-champions-2024-gf",
				Region:      "PACIFIC#INTERNATIONAL",
				StartTime:   "2024-08-25T07:00:00+0000",
				TeamAway:    "Team Heretics",
				TeamHome:    "EDward Gaming",
				TTL:         1724630400,
			},
		},
		{
			name: "lack of information",
			match: domain.Match{
				Id:               999999,
				Name:             "Group Stage",
				StartDate:        "2024-01-01",
				StartTime:        "2024-01-01T00:00:00+0000",
				BestOf:           0, // default to 0
				Teams:            []domain.Team{{Name: "TBD"}, {Name: "TBD"}},
				PagePath:         "/999999/match-with-lack-of-information",
				EventId:          "9999",
				EventName:        "Empty Event",
				EventCountryFlag: "un",
			},
			want: dto.Match{
				MatchID:     999999,
				BestOf:      0,
				EndTime:     "2024-01-01T03:00:00+0000",
				EventDetail: "Group Stage",
				EventName:   "Empty Event",
				MatchURI:    "https://vlr.gg/999999/match-with-lack-of-information",
				Region:      "",
				StartTime:   "2024-01-01T00:00:00+0000",
				TeamAway:    "TBD",
				TeamHome:    "TBD",
				TTL:         1704121200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchToDTO(tt.match)
			if got != tt.want {
				t.Errorf("MatchToDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapFlagToRegion(t *testing.T) {
	tests := []struct {
		name        string
		countryFlag string
		want        string
	}{
		{"Global Region - 1", "de", "EMEA"},
		{"Global Region - 2", "cn", "CHINA"},
		{"Local Region - 1", "br", "AMERICAS"},
		{"Global Flag", "un", ""},
		{"Empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapFlagToRegion(tt.countryFlag)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("mapFlagToRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEstimateRegionByEvent(t *testing.T) {
	tests := []struct {
		name      string
		eventName string
		want      string
	}{
		{
			name:      "Found - Sub Area",
			eventName: "Road to VCL MENA: North Africa and Levant #1",
			want:      "EMEA",
		},
		{
			name:      "Found - Organizer",
			eventName: "Global Esports' 2025 Fight Night #3",
			want:      "PACIFIC",
		},
		{
			name:      "Not Found",
			eventName: "Some Unknown Event",
			want:      "",
		},
		{
			name:      "Empty",
			eventName: "",
			want:      "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EstimateRegionByEvent(tt.eventName)
			if got != tt.want {
				t.Errorf("EstimateRegionByEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
