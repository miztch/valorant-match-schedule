package infrastructure

import (
	"context"

	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// GoogleCalendarServiceImpl implements domain.GoogleCalendarService
type GoogleCalendarServiceImpl struct {
	client *GoogleCalendarClient
}

// NewGoogleCalendarService creates a new Google Calendar service
func NewGoogleCalendarService(ctx context.Context) (domain.GoogleCalendarService, error) {
	client, err := NewGoogleCalendarClient(ctx)
	if err != nil {
		return nil, err
	}
	return &GoogleCalendarServiceImpl{
		client: client,
	}, nil
}

// CreateEvent creates a calendar event with a specific ID
func (s *GoogleCalendarServiceImpl) CreateEvent(ctx context.Context, calendarID string, matchItem *dto.MatchForStream, eventID string) (string, error) {
	event, err := s.client.CreateEventFromMatchItemWithID(ctx, calendarID, matchItem, eventID)
	if err != nil {
		return "", err
	}

	return event.Id, nil
}

// UpdateEvent updates an existing calendar event
func (s *GoogleCalendarServiceImpl) UpdateEvent(ctx context.Context, calendarID string, matchItem *dto.MatchForStream, eventID string) (string, error) {
	event, err := s.client.UpdateEventFromMatchItemWithID(ctx, calendarID, matchItem, eventID)
	if err != nil {
		return "", err
	}

	return event.Id, nil
}

// DeleteEvent deletes a calendar event
func (s *GoogleCalendarServiceImpl) DeleteEvent(ctx context.Context, calendarID, eventID string) error {
	err := s.client.DeleteEvent(ctx, calendarID, eventID)
	if err != nil {
		return err
	}

	return nil
}
