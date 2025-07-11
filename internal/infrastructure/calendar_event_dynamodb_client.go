package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// Create DynamoDB client for calendar event registration
func NewCalendarEventDynamoDBClient(ctx context.Context) (*DynamoDBClient, error) {
	tableName := os.Getenv("CALENDAR_EVENT_TABLE")
	if tableName == "" {
		return nil, fmt.Errorf("CALENDAR_EVENT_TABLE is not set")
	}
	return NewDynamoDBClient(ctx, tableName)
}

// Save calendar event to DynamoDB
func WriteCalendarEvent(ctx context.Context, client *DynamoDBClient, event dto.CalendarEvent) error {
	return PutItems(ctx, client, []dto.CalendarEvent{event})
}

// Check calendar event existence (search by match_id, calendar_id)
// Return values: (exists, event_id, error)
func CalendarEventExists(ctx context.Context, client *DynamoDBClient, matchID int, calendarID string) (bool, string, error) {
	key := map[string]types.AttributeValue{
		"match_id":    &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", matchID)},
		"calendar_id": &types.AttributeValueMemberS{Value: calendarID},
	}
	out, err := client.GetItem(ctx, key)
	if err != nil {
		return false, "", err
	}
	if len(out.Item) == 0 {
		return false, "", nil
	}
	var ce dto.CalendarEvent
	err = attributevalue.UnmarshalMap(out.Item, &ce)
	if err != nil {
		return true, "", err
	}
	return true, ce.EventID, nil
}

// Get all calendar events for specified match_id
func ListCalendarEventsForMatch(ctx context.Context, client *DynamoDBClient, matchID int) ([]dto.CalendarEvent, error) {
	input := &dynamodb.QueryInput{
		TableName: &client.TableName,
		KeyConditions: map[string]types.Condition{
			"match_id": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: fmt.Sprintf("%d", matchID)},
				},
			},
		},
	}
	out, err := client.client.Query(ctx, input)
	if err != nil {
		return nil, err
	}
	var events []dto.CalendarEvent
	err = attributevalue.UnmarshalListOfMaps(out.Items, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Delete calendar event from DynamoDB
func DeleteCalendarEvent(ctx context.Context, client *DynamoDBClient, matchID int, calendarID string) error {
	key := map[string]types.AttributeValue{
		"match_id":    &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", matchID)},
		"calendar_id": &types.AttributeValueMemberS{Value: calendarID},
	}
	_, err := client.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &client.TableName,
		Key:       key,
	})
	return err
}
