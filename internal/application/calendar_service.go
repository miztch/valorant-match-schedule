package application

import (
	"context"
	"log/slog"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
	"github.com/miztch/valorant-match-schedule/internal/infrastructure"
)

type CalendarService struct {
	calendarEventRepo domain.CalendarEventRepository
	googleCalendarSvc domain.GoogleCalendarService
	streamParser      domain.DynamoDBStreamParser
	calendarIDs       map[string]string
}

// NewCalendarService creates a new CalendarService with dependencies injected
func NewCalendarService(
	calendarEventRepo domain.CalendarEventRepository,
	googleCalendarSvc domain.GoogleCalendarService,
	streamParser domain.DynamoDBStreamParser,
	calendarIDs map[string]string,
) *CalendarService {
	return &CalendarService{
		calendarEventRepo: calendarEventRepo,
		googleCalendarSvc: googleCalendarSvc,
		streamParser:      streamParser,
		calendarIDs:       calendarIDs,
	}
}

// NewCalendarServiceFromEnv creates a service with default implementations from environment
func NewCalendarServiceFromEnv(ctx context.Context) (*CalendarService, error) {
	calendarEventRepo, err := infrastructure.NewCalendarEventRepository(ctx)
	if err != nil {
		return nil, err
	}

	googleCalendarSvc, err := infrastructure.NewGoogleCalendarService(ctx)
	if err != nil {
		return nil, err
	}

	streamParser := infrastructure.NewDynamoDBStreamParser()
	calendarIDs := infrastructure.LoadCalendarIDsFromEnv()

	return NewCalendarService(calendarEventRepo, googleCalendarSvc, streamParser, calendarIDs), nil
}

// HandleDynamoDBStream processes DynamoDB Stream events to create, update, or delete Google Calendar events.
func (svc *CalendarService) HandleDynamoDBStream(ctx context.Context, event events.DynamoDBEvent) error {
	for _, record := range event.Records {
		if record.EventName != "INSERT" && record.EventName != "MODIFY" {
			continue
		}
		svc.processRecord(ctx, record)
	}
	return nil
}

// processRecord processes a single DynamoDB event record to create, update and delete calendar events.
func (svc *CalendarService) processRecord(ctx context.Context, record events.DynamoDBEventRecord) {
	matchItem, err := svc.streamParser.ParseMatchFromStream(record.Change.NewImage)
	if err != nil {
		slog.Error("Failed to unmarshal DynamoDB stream image", "error", err.Error())
		return
	}

	regionStr, ok := matchItem.GetRegion()
	if !ok || regionStr == "" {
		slog.Warn("item.region is empty, skipping calendar registration")
		return
	}

	matchID, ok := matchItem.GetMatchID()
	if !ok {
		slog.Error("Missing match_id field in DynamoDB stream item")
		return
	}

	regionParts := strings.Split(regionStr, "#")
	usedCalendars := make(map[string]bool)

	// Remove obsolete calendar events
	svc.removeObsoleteCalendarEvents(ctx, matchID, regionParts)
	// Create/update necessary calendar events
	svc.upsertCalendarEventsFromMatchItem(ctx, matchItem, regionParts, usedCalendars)
}

func (svc *CalendarService) removeObsoleteCalendarEvents(ctx context.Context, matchID int, regionParts []string) {
	existingRegions, err := svc.calendarEventRepo.ListByMatch(ctx, matchID)
	if err != nil {
		slog.Error("Failed to list calendar events for match", "error", err.Error())
		return
	}
	regionSet := make(map[string]struct{})
	for _, r := range regionParts {
		regionSet[r] = struct{}{}
	}
	for _, ce := range existingRegions {
		if _, ok := regionSet[ce.Region]; !ok {
			// Log before deleting
			slog.Info("Deleting Google Calendar event",
				"calendarId", ce.CalendarID,
				"eventId", ce.EventID,
			)

			// Delete from calendar
			err := svc.googleCalendarSvc.DeleteEvent(ctx, ce.CalendarID, ce.EventID)
			if err != nil {
				slog.Error("Failed to delete Google Calendar event",
					"error", err.Error(),
					"calendarId", ce.CalendarID,
					"eventId", ce.EventID,
				)
			}
			// Also delete from table
			err = svc.calendarEventRepo.Delete(ctx, ce.MatchID, ce.CalendarID)
			if err != nil {
				slog.Error("Failed to delete calendar event from DynamoDB", "calendarID", ce.CalendarID, "error", err.Error())
			}
		}
	}
}

func (svc *CalendarService) upsertCalendarEventsFromMatchItem(ctx context.Context, matchItem *dto.MatchForStream, regionParts []string, usedCalendars map[string]bool) {
	for _, region := range regionParts {
		calID, ok := svc.calendarIDs[region]
		if !ok || calID == "" {
			slog.Warn("No calendar ID for region", "region", region)
			continue
		}
		if usedCalendars[calID] {
			continue
		}

		matchID, ok := matchItem.GetMatchID()
		if !ok {
			slog.Error("Missing match_id field in match item")
			continue
		}

		genEventID := svc.streamParser.GenerateEventID(matchID)
		var eventID string
		exists, oldEventID, err := svc.calendarEventRepo.Exists(ctx, matchID, calID)
		if err != nil {
			slog.Error("Failed to check calendar event existence", "error", err.Error())
			continue
		}
		if exists {
			eventID = oldEventID

			// Log before updating
			matchID, _ := matchItem.GetMatchID()
			region, _ := matchItem.GetRegion()
			eventName, _ := matchItem.GetEventName()
			teamHome, _ := matchItem.GetTeamHome()
			teamAway, _ := matchItem.GetTeamAway()
			startTime, _ := matchItem.GetStartTime()

			slog.Info("Updating Google Calendar event",
				"calendarId", calID,
				"region", region,
				"eventId", eventID,
				"matchId", matchID,
				"eventName", eventName,
				"teams", teamHome+" vs "+teamAway,
				"startTime", startTime,
			)

			createdEventID, err := svc.googleCalendarSvc.UpdateEvent(ctx, calID, matchItem, eventID)
			if err != nil {
				slog.Error("Failed to update Google Calendar event",
					"error", err.Error(),
					"calendarId", calID,
					"region", region,
					"eventId", eventID,
					"matchId", matchID,
				)
				continue
			}
			eventID = createdEventID
		} else {
			// Log before creating
			matchID, _ := matchItem.GetMatchID()
			region, _ := matchItem.GetRegion()
			eventName, _ := matchItem.GetEventName()
			teamHome, _ := matchItem.GetTeamHome()
			teamAway, _ := matchItem.GetTeamAway()
			startTime, _ := matchItem.GetStartTime()

			slog.Info("Creating Google Calendar event",
				"calendarId", calID,
				"region", region,
				"eventId", genEventID,
				"matchId", matchID,
				"eventName", eventName,
				"teams", teamHome+" vs "+teamAway,
				"startTime", startTime,
			)

			createdEventID, err := svc.googleCalendarSvc.CreateEvent(ctx, calID, matchItem, genEventID)
			if err != nil {
				slog.Error("Failed to create Google Calendar event",
					"error", err.Error(),
					"calendarId", calID,
					"region", region,
					"eventId", genEventID,
					"matchId", matchID,
				)
				continue
			}
			eventID = createdEventID
		}

		ttl, ok := matchItem.GetTTL()
		if !ok {
			slog.Error("Missing ttl field in match item")
			continue
		}
		ttl += 30 * 24 * 60 * 60 // Add 30 days

		calendarEvent := dto.CalendarEvent{
			MatchID:    matchID,
			CalendarID: calID,
			EventID:    eventID,
			TTL:        ttl,
			Region:     region,
		}
		err = svc.calendarEventRepo.Save(ctx, calendarEvent)
		if err != nil {
			slog.Error("Failed to write calendar event to DynamoDB", "error", err.Error())
		}
		usedCalendars[calID] = true
	}
}
