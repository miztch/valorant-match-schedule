package infrastructure

import (
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// DynamoDBStreamParserImpl implements domain.DynamoDBStreamParser
type DynamoDBStreamParserImpl struct{}

// NewDynamoDBStreamParser creates a new DynamoDB stream parser
func NewDynamoDBStreamParser() domain.DynamoDBStreamParser {
	return &DynamoDBStreamParserImpl{}
}

// ParseMatchFromStream parses DynamoDB stream image to MatchForStream
func (p *DynamoDBStreamParserImpl) ParseMatchFromStream(image map[string]events.DynamoDBAttributeValue) (*dto.MatchForStream, error) {
	return UnmarshalDynamoDBStreamImageToMatchForStream(image)
}

// GenerateEventID generates a consistent event ID for Google Calendar
func (p *DynamoDBStreamParserImpl) GenerateEventID(matchID int) string {
	return GenerateGoogleCalendarEventID(matchID)
}

// Helper function to generate Google Calendar event ID
// The generated ID is in the format "match<matchID><base32hex>"
func GenerateGoogleCalendarEventID(matchID int) string {
	// Generate a deterministic suffix based on a UUID
	u := uuid.New()
	enc := base32.NewEncoding("0123456789abcdefghijklmnopqrstuv").WithPadding(base32.NoPadding)
	base32id := strings.ToLower(enc.EncodeToString(u[:8]))
	return fmt.Sprintf("match%d%s", matchID, base32id)
}
