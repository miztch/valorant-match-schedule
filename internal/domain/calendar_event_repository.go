package domain

import (
	"context"

	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// CalendarEventRepository defines the interface for calendar event persistence operations
type CalendarEventRepository interface {
	// Create or update a calendar event in the data store
	Save(ctx context.Context, event dto.CalendarEvent) error

	// Check if a calendar event exists for the given match and calendar
	// Returns: (exists, eventID, error)
	Exists(ctx context.Context, matchID int, calendarID string) (bool, string, error)

	// List all calendar events for a specific match
	ListByMatch(ctx context.Context, matchID int) ([]dto.CalendarEvent, error)

	// Delete a calendar event by match and calendar ID
	Delete(ctx context.Context, matchID int, calendarID string) error
}
