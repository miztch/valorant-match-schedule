package domain

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

// Match represents a match (for business logic only, no DTO tags)
type Match struct {
	Id               int
	Name             string
	StartDate        string
	StartTime        string
	BestOf           int
	Teams            []Team
	PagePath         string
	EventName        string
	EventCountryFlag string
}

// Team represents a team
// DTO tags removed
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
		EventName:        e.Name,
		EventCountryFlag: e.CountryFlag,
	}
}
