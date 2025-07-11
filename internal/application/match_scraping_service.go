package application

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/miztch/valorant-match-schedule/internal/domain"
)

type MatchService struct {
	matchRepo domain.MatchRepository
	eventRepo domain.EventRepository
}

func NewMatchService(matchRepo domain.MatchRepository, eventRepo domain.EventRepository) *MatchService {
	return &MatchService{
		matchRepo: matchRepo,
		eventRepo: eventRepo,
	}
}

func (svc *MatchService) FetchMatches(page int) ([]domain.Match, error) {
	slog.Info("Starting to fetch matches", "page", page)

	// Implement match fetching logic
	matchURLs, err := svc.matchRepo.GetMatchURLList(page)
	if err != nil {
		slog.Error("Failed to get match URLs", "error", err.Error(), "page", page)
		return nil, fmt.Errorf("failed to get match urls: %w", err)
	}

	slog.Info("Retrieved match URLs", "count", len(matchURLs), "page", page)

	var matches []domain.Match
	for i, matchURL := range matchURLs {
		m, err := svc.matchRepo.ScrapeMatch(matchURL)
		if err != nil {
			slog.Error("Failed to scrape match", "error", err.Error(), "matchURL", matchURL, "index", i)
			return nil, fmt.Errorf("failed to get match: %w", err)
		}

		// Check if the scraped match is empty
		if domain.IsEmptyVlrMatch(m) {
			slog.Warn("Skipping empty match", "matchURL", matchURL, "index", i)
			continue
		}

		slog.Info("Successfully scraped match", "matchId", m.Id, "matchName", m.Name, "matchURL", matchURL)

		e, err := svc.eventRepo.GetEvent(m.EventPagePath)
		if err != nil {
			slog.Error("Failed to get event", "error", err.Error(), "eventPath", m.EventPagePath, "matchId", m.Id)
			return nil, fmt.Errorf("failed to get event: %w", err)
		}

		match := domain.NewMatch(m, e)
		matches = append(matches, match)
	}

	slog.Info("Successfully fetched matches", "totalCount", len(matches), "page", page)
	return matches, nil
}

func (svc *MatchService) WriteMatches(ctx context.Context, matches []domain.Match) error {
	slog.Info("Writing matches to repository", "count", len(matches))

	// Implement match writing logic
	err := svc.matchRepo.WriteMatches(ctx, matches)
	if err != nil {
		slog.Error("Failed to write matches", "error", err.Error(), "count", len(matches))
		return fmt.Errorf("failed to write matches: %w", err)
	}

	slog.Info("Successfully wrote matches", "count", len(matches))
	return nil
}
