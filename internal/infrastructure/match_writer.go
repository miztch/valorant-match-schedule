package infrastructure

import (
	"context"
	"log/slog"

	"github.com/miztch/valorant-match-schedule/internal/domain"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// MatchDynamoDBWriter implements domain.MatchWriter for DynamoDB storage
type MatchDynamoDBWriter struct {
	client *DynamoDBClient
}

// NewMatchWriter creates a new match writer for DynamoDB
func NewMatchWriter(ctx context.Context) (domain.MatchWriter, error) {
	client, err := NewMatchDynamoDBClient(ctx)
	if err != nil {
		return nil, err
	}
	return &MatchDynamoDBWriter{
		client: client,
	}, nil
}

// WriteMatches writes multiple matches to DynamoDB
func (w *MatchDynamoDBWriter) WriteMatches(ctx context.Context, matches []dto.Match) error {
	slog.Info("Writing matches to DynamoDB", "count", len(matches))

	// Log each match for debugging
	for i, match := range matches {
		slog.Info("Match to be written to DynamoDB",
			"index", i,
			"matchId", match.MatchID,
			"eventName", match.EventName,
			"region", match.Region,
			"startTime", match.StartTime,
			"teams", match.TeamHome+" vs "+match.TeamAway,
		)
	}

	err := WriteMatches(ctx, w.client, matches)
	if err != nil {
		slog.Error("Failed to write matches to DynamoDB", "error", err.Error(), "count", len(matches))
		return err
	}

	slog.Info("Successfully wrote matches to DynamoDB", "count", len(matches))
	return nil
}
