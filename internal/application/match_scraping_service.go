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

func (svc *MatchService) FetchMatches(page int) ([]domain.Match, int, error) {
	slog.Info("Starting to fetch matches", "page", page)

	matchURLs, err := svc.matchRepo.GetMatchURLList(page)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get match urls: %w", err)
	}

	slog.Info("Retrieved match URLs", "count", len(matchURLs), "page", page)

	var matches []domain.Match
	for i, matchURL := range matchURLs {
		m, err := svc.matchRepo.ScrapeMatch(matchURL)
		if err != nil {
			slog.Warn("Failed to scrape match, skipping", "error", err.Error(), "matchURL", matchURL, "index", i)
			continue
		}

		if domain.IsEmptyVlrMatch(m) {
			slog.Warn("Skipping empty match", "matchURL", matchURL, "index", i)
			continue
		}

		slog.Info("Successfully scraped match", "matchId", m.Id, "matchName", m.Name, "matchURL", matchURL)

		e, err := svc.eventRepo.GetEvent(m.EventPagePath)
		if err != nil {
			slog.Warn("Failed to get event, skipping", "error", err.Error(), "eventPath", m.EventPagePath, "matchId", m.Id)
			continue
		}

		match := domain.NewMatch(m, e)

		// avoid matchId duplication
		for _, existingMatch := range matches {
			if existingMatch.Id == match.Id {
				slog.Warn("Duplicate match ID found, skipping", "matchId", match.Id, "matchURL", matchURL)
				continue
			}
		}
		matches = append(matches, match)
	}

	slog.Info("Successfully fetched matches", "totalCount", len(matches), "urlCount", len(matchURLs), "page", page)
	return matches, len(matchURLs), nil
}

func (svc *MatchService) WriteMatches(ctx context.Context, matches []domain.Match) error {
	slog.Info("Writing matches to repository", "count", len(matches))

	err := svc.matchRepo.WriteMatches(ctx, matches)
	if err != nil {
		return fmt.Errorf("failed to write matches: %w", err)
	}

	slog.Info("Successfully wrote matches", "count", len(matches))
	return nil
}
