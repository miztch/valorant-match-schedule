package application

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// MatchRegistrationService handles registration of matches from SQS events
type MatchRegistrationService struct {
	matchWriter domain.MatchWriter
}

// NewMatchRegistrationService creates a new match registration service
func NewMatchRegistrationService(matchWriter domain.MatchWriter) *MatchRegistrationService {
	return &MatchRegistrationService{
		matchWriter: matchWriter,
	}
}

// ProcessSQSEvent processes SQS events containing match data
func (s *MatchRegistrationService) ProcessSQSEvent(ctx context.Context, sqsEvent events.SQSEvent) error {
	var matches []dto.Match
	var processedRecords []string

	for i, record := range sqsEvent.Records {
		slog.Info("Processing SQS record",
			"recordIndex", i,
			"messageId", record.MessageId,
			"eventSource", record.EventSource,
			"eventSourceARN", record.EventSourceARN,
		)

		// Parse the match data from the message body
		var match domain.Match
		if err := json.Unmarshal([]byte(record.Body), &match); err != nil {
			slog.Error("Failed to unmarshal match data",
				"error", err.Error(),
				"messageId", record.MessageId,
				"body", record.Body,
			)
			continue
		}

		// Log the match details
		slog.Info("Match data received",
			"matchId", match.Id,
			"matchName", match.Name,
			"eventName", match.EventName,
			"startDate", match.StartDate,
			"startTime", match.StartTime,
			"bestOf", match.BestOf,
			"teams", match.Teams,
			"eventCountryFlag", match.EventCountryFlag,
			"pagePath", match.PagePath,
		)

		// Output Match JSON
		matchDTO := MatchToDTO(match)
		matchJSON, _ := json.Marshal(matchDTO)
		slog.Info("Match JSON", "json", string(matchJSON))

		// Add Match to slice
		matches = append(matches, matchDTO)
		processedRecords = append(processedRecords, record.MessageId)
	}

	// Batch write to DynamoDB only if there are processed data
	if len(matches) > 0 {
		err := s.matchWriter.WriteMatches(ctx, matches)
		if err != nil {
			slog.Error("Failed to write Matches to DynamoDB",
				"error", err.Error(),
				"count", len(matches),
				"processedRecords", processedRecords,
			)
			return err
		}

		slog.Info("Successfully saved Matches to DynamoDB",
			"count", len(matches),
			"processedRecords", processedRecords,
		)
	}

	slog.Info("Completed processing SQS event",
		"totalRecords", len(sqsEvent.Records),
		"successfullyProcessed", len(matches),
	)
	return nil
}
