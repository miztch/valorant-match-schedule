package dto

type CalendarEvent struct {
	MatchID    int    `json:"match_id" dynamodbav:"match_id"`
	CalendarID string `json:"calendar_id" dynamodbav:"calendar_id"`
	EventID    string `json:"event_id" dynamodbav:"event_id"`
	TTL        int64  `json:"ttl" dynamodbav:"ttl"`
	Region     string `json:"region" dynamodbav:"region"`
}
