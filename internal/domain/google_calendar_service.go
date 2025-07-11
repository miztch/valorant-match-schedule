package domain

import (
	"context"

	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// GoogleCalendarService defines the interface for Google Calendar operations
type GoogleCalendarService interface {
	// Create a calendar event with a specific ID
	CreateEvent(ctx context.Context, calendarID string, matchItem *dto.MatchForStream, eventID string) (string, error)

	// Update an existing calendar event
	UpdateEvent(ctx context.Context, calendarID string, matchItem *dto.MatchForStream, eventID string) (string, error)

	// Delete a calendar event
	DeleteEvent(ctx context.Context, calendarID, eventID string) error
}
