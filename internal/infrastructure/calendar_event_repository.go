package infrastructure

import (
	"context"

	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// CalendarEventRepositoryImpl implements domain.CalendarEventRepository
type CalendarEventRepositoryImpl struct {
	client *DynamoDBClient
}

// NewCalendarEventRepository creates a new calendar event repository
func NewCalendarEventRepository(ctx context.Context) (domain.CalendarEventRepository, error) {
	client, err := NewCalendarEventDynamoDBClient(ctx)
	if err != nil {
		return nil, err
	}
	return &CalendarEventRepositoryImpl{
		client: client,
	}, nil
}

// Save stores a calendar event in the data store
func (r *CalendarEventRepositoryImpl) Save(ctx context.Context, event dto.CalendarEvent) error {
	return WriteCalendarEvent(ctx, r.client, event)
}

// Exists checks if a calendar event exists for the given match and calendar
func (r *CalendarEventRepositoryImpl) Exists(ctx context.Context, matchID int, calendarID string) (bool, string, error) {
	return CalendarEventExists(ctx, r.client, matchID, calendarID)
}

// ListByMatch returns all calendar events for a specific match
func (r *CalendarEventRepositoryImpl) ListByMatch(ctx context.Context, matchID int) ([]dto.CalendarEvent, error) {
	return ListCalendarEventsForMatch(ctx, r.client, matchID)
}

// Delete removes a calendar event by match and calendar ID
func (r *CalendarEventRepositoryImpl) Delete(ctx context.Context, matchID int, calendarID string) error {
	return DeleteCalendarEvent(ctx, r.client, matchID, calendarID)
}
