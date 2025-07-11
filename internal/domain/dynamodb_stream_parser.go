package domain

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// DynamoDBStreamParser defines the interface for parsing DynamoDB stream events
type DynamoDBStreamParser interface {
	// Parse DynamoDB stream image to MatchForStream
	ParseMatchFromStream(image map[string]events.DynamoDBAttributeValue) (*dto.MatchForStream, error)

	// Generate a consistent event ID for Google Calendar
	GenerateEventID(matchID int) string
}
