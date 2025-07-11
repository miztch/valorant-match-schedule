package dto

// Match represents a match item in the system
// Used for both writing to and reading from DynamoDB
type Match struct {
	MatchID     int    `json:"match_id" dynamodbav:"match_id"`
	BestOf      int    `json:"best_of" dynamodbav:"best_of"`
	EndTime     string `json:"end_time" dynamodbav:"end_time"`
	EventDetail string `json:"event_detail" dynamodbav:"event_detail"`
	EventName   string `json:"event_name" dynamodbav:"event_name"`
	MatchURI    string `json:"match_uri" dynamodbav:"match_uri"`
	Region      string `json:"region" dynamodbav:"region"`
	StartTime   string `json:"start_time" dynamodbav:"start_time"`
	TeamAway    string `json:"team_away" dynamodbav:"team_away"`
	TeamHome    string `json:"team_home" dynamodbav:"team_home"`
	TTL         int64  `json:"ttl" dynamodbav:"ttl"`
}

// MatchForStream represents a match item from DynamoDB stream with nullable fields
type MatchForStream struct {
	MatchID     *int    `dynamodbav:"match_id"`
	BestOf      *int    `dynamodbav:"best_of"`
	EndTime     *string `dynamodbav:"end_time"`
	EventDetail *string `dynamodbav:"event_detail"`
	EventName   *string `dynamodbav:"event_name"`
	MatchURI    *string `dynamodbav:"match_uri"`
	Region      *string `dynamodbav:"region"`
	StartTime   *string `dynamodbav:"start_time"`
	TeamAway    *string `dynamodbav:"team_away"`
	TeamHome    *string `dynamodbav:"team_home"`
	TTL         *int64  `dynamodbav:"ttl"`
}

// ToMatch converts MatchForStream to Match with default values for nil fields
func (m *MatchForStream) ToMatch() Match {
	return Match{
		MatchID:     getIntValue(m.MatchID),
		BestOf:      getIntValue(m.BestOf),
		EndTime:     getStringValue(m.EndTime),
		EventDetail: getStringValue(m.EventDetail),
		EventName:   getStringValue(m.EventName),
		MatchURI:    getStringValue(m.MatchURI),
		Region:      getStringValue(m.Region),
		StartTime:   getStringValue(m.StartTime),
		TeamAway:    getStringValue(m.TeamAway),
		TeamHome:    getStringValue(m.TeamHome),
		TTL:         getInt64Value(m.TTL),
	}
}

// Safe getter methods for MatchForStream
func (m *MatchForStream) GetMatchID() (int, bool) {
	if m.MatchID == nil {
		return 0, false
	}
	return *m.MatchID, true
}

func (m *MatchForStream) GetRegion() (string, bool) {
	if m.Region == nil {
		return "", false
	}
	return *m.Region, true
}

func (m *MatchForStream) GetTTL() (int64, bool) {
	if m.TTL == nil {
		return 0, false
	}
	return *m.TTL, true
}

func (m *MatchForStream) GetEventName() (string, bool) {
	if m.EventName == nil {
		return "", false
	}
	return *m.EventName, true
}

func (m *MatchForStream) GetEventDetail() (string, bool) {
	if m.EventDetail == nil {
		return "", false
	}
	return *m.EventDetail, true
}

func (m *MatchForStream) GetTeamHome() (string, bool) {
	if m.TeamHome == nil {
		return "", false
	}
	return *m.TeamHome, true
}

func (m *MatchForStream) GetTeamAway() (string, bool) {
	if m.TeamAway == nil {
		return "", false
	}
	return *m.TeamAway, true
}

func (m *MatchForStream) GetStartTime() (string, bool) {
	if m.StartTime == nil {
		return "", false
	}
	return *m.StartTime, true
}

func (m *MatchForStream) GetEndTime() (string, bool) {
	if m.EndTime == nil {
		return "", false
	}
	return *m.EndTime, true
}

func (m *MatchForStream) GetMatchURI() (string, bool) {
	if m.MatchURI == nil {
		return "", false
	}
	return *m.MatchURI, true
}

// Helper functions for safe conversion
func getIntValue(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func getInt64Value(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
