package infrastructure

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/miztch/valorant-match-schedule/internal/dto"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarClient struct {
	service *calendar.Service
}

// Create Google Calendar API client
func NewGoogleCalendarClient(ctx context.Context) (*GoogleCalendarClient, error) {
	// Get service account key from SSM Parameter Store
	keyData, err := GetServiceAccountKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get service account key from SSM: %w", err)
	}

	// Create Google API client with service account key
	creds, err := google.CredentialsFromJSON(ctx, keyData, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials from JSON: %w", err)
	}

	service, err := calendar.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to create calendar service: %w", err)
	}
	return &GoogleCalendarClient{service: service}, nil
}

// Convert DynamoDB Stream NewImage (events.DynamoDBAttributeValue) to MatchForStream struct
func UnmarshalDynamoDBStreamImageToMatchForStream(image map[string]events.DynamoDBAttributeValue) (*dto.MatchForStream, error) {
	// Convert events.DynamoDBAttributeValue to types.AttributeValue
	convertedImage := make(map[string]types.AttributeValue)
	for k, v := range image {
		convertedImage[k] = convertEventAttributeValueToSDKAttributeValue(v)
	}

	var item dto.MatchForStream
	err := attributevalue.UnmarshalMap(convertedImage, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB stream image: %w", err)
	}

	return &item, nil
}

// Helper to convert events.DynamoDBAttributeValue to types.AttributeValue
func convertEventAttributeValueToSDKAttributeValue(eventValue events.DynamoDBAttributeValue) types.AttributeValue {
	switch eventValue.DataType() {
	case events.DataTypeString:
		return &types.AttributeValueMemberS{Value: eventValue.String()}
	case events.DataTypeNumber:
		return &types.AttributeValueMemberN{Value: eventValue.Number()}
	case events.DataTypeBoolean:
		return &types.AttributeValueMemberBOOL{Value: eventValue.Boolean()}
	case events.DataTypeBinary:
		return &types.AttributeValueMemberB{Value: eventValue.Binary()}
	case events.DataTypeList:
		// List case requires more complex conversion, omitted for current use case
		return &types.AttributeValueMemberS{Value: "LIST_NOT_SUPPORTED"}
	case events.DataTypeMap:
		// Map case requires more complex conversion, omitted for current use case
		return &types.AttributeValueMemberS{Value: "MAP_NOT_SUPPORTED"}
	case events.DataTypeNull:
		return &types.AttributeValueMemberNULL{Value: true}
	default:
		return &types.AttributeValueMemberS{Value: eventValue.String()}
	}
}

// Create Google Calendar event (struct version)
func (c *GoogleCalendarClient) CreateEventFromMatchItem(ctx context.Context, calendarID string, item *dto.MatchForStream) (*calendar.Event, error) {
	eventName, _ := item.GetEventName()
	eventDetail, _ := item.GetEventDetail()
	teamHome, _ := item.GetTeamHome()
	teamAway, _ := item.GetTeamAway()
	startTime, _ := item.GetStartTime()
	endTime, _ := item.GetEndTime()
	matchURI, _ := item.GetMatchURI()

	event := &calendar.Event{
		Summary:     fmt.Sprintf("%s - %s | %s - %s", teamHome, teamAway, eventName, eventDetail),
		Description: matchURI,
		Start:       &calendar.EventDateTime{DateTime: startTime, TimeZone: "Etc/UTC"},
		End:         &calendar.EventDateTime{DateTime: endTime, TimeZone: "Etc/UTC"},
	}
	created, err := c.service.Events.Insert(calendarID, event).Context(ctx).Do()
	return created, err
}

// Create Google Calendar event (struct version, with ID specification)
func (c *GoogleCalendarClient) CreateEventFromMatchItemWithID(ctx context.Context, calendarID string, item *dto.MatchForStream, eventID string) (*calendar.Event, error) {
	eventName, _ := item.GetEventName()
	eventDetail, _ := item.GetEventDetail()
	teamHome, _ := item.GetTeamHome()
	teamAway, _ := item.GetTeamAway()
	startTime, _ := item.GetStartTime()
	endTime, _ := item.GetEndTime()
	matchURI, _ := item.GetMatchURI()

	event := &calendar.Event{
		Id:          eventID,
		Summary:     fmt.Sprintf("%s - %s | %s - %s", teamHome, teamAway, eventName, eventDetail),
		Description: matchURI,
		Start:       &calendar.EventDateTime{DateTime: startTime, TimeZone: "Etc/UTC"},
		End:         &calendar.EventDateTime{DateTime: endTime, TimeZone: "Etc/UTC"},
	}
	created, err := c.service.Events.Insert(calendarID, event).Context(ctx).Do()
	return created, err
}

// Update Google Calendar event (struct version, with ID specification)
func (c *GoogleCalendarClient) UpdateEventFromMatchItemWithID(ctx context.Context, calendarID string, item *dto.MatchForStream, eventID string) (*calendar.Event, error) {
	eventName, _ := item.GetEventName()
	eventDetail, _ := item.GetEventDetail()
	teamHome, _ := item.GetTeamHome()
	teamAway, _ := item.GetTeamAway()
	startTime, _ := item.GetStartTime()
	endTime, _ := item.GetEndTime()
	matchURI, _ := item.GetMatchURI()

	event := &calendar.Event{
		Summary:     fmt.Sprintf("%s - %s | %s - %s", teamHome, teamAway, eventName, eventDetail),
		Description: matchURI,
		Start:       &calendar.EventDateTime{DateTime: startTime, TimeZone: "Etc/UTC"},
		End:         &calendar.EventDateTime{DateTime: endTime, TimeZone: "Etc/UTC"},
	}
	updated, err := c.service.Events.Update(calendarID, eventID, event).Context(ctx).Do()
	return updated, err
}

// Delete Google Calendar event
func (c *GoogleCalendarClient) DeleteEvent(ctx context.Context, calendarID, eventID string) error {
	return c.service.Events.Delete(calendarID, eventID).Context(ctx).Do()
}

// Load calendar IDs from environment variables as region:value pairs
func LoadCalendarIDsFromEnv() map[string]string {
	ids := make(map[string]string)
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "CALENDAR_ID_") {
			parts := strings.SplitN(e, "=", 2)
			if len(parts) == 2 && parts[1] != "" {
				region := strings.TrimPrefix(parts[0], "CALENDAR_ID_")
				ids[region] = parts[1]
			}
		}
	}
	return ids
}
