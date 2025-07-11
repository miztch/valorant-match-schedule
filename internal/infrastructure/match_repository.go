package infrastructure

import (
	"context"
	"fmt"

	"github.com/miztch/valorant-match-schedule/internal/domain"
)

// MatchRepository implementation for scraping and writing matches
type matchRepositoryImpl struct {
	vlrGGScraper *VlrGGScraper
	sqsClient    *SQSClient
}

// NewMatchRepository creates a new match repository
func NewMatchRepository(vlrGGScraper *VlrGGScraper, sqsClient *SQSClient) *matchRepositoryImpl {
	return &matchRepositoryImpl{
		vlrGGScraper: vlrGGScraper,
		sqsClient:    sqsClient,
	}
}

// ScrapeMatch scrapes a match
func (r *matchRepositoryImpl) ScrapeMatch(matchUrlPath string) (domain.VlrMatch, error) {
	match, err := r.vlrGGScraper.scrapeMatch(matchUrlPath)
	if err != nil {
		return domain.VlrMatch{}, fmt.Errorf("failed to scrape match: %w", err)
	}
	return match, nil
}

// GetMatchURLList gets a list of match URLs
func (r *matchRepositoryImpl) GetMatchURLList(pageNumber int) ([]string, error) {
	matchURLs, err := r.vlrGGScraper.getMatchURLList(pageNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get match urls: %w", err)
	}

	return matchURLs, nil
}

// WriteMatches writes matches to SQS
func (r *matchRepositoryImpl) WriteMatches(ctx context.Context, matches []domain.Match) error {
	err := r.sqsClient.WriteMatches(ctx, matches)
	if err != nil {
		return fmt.Errorf("failed to write matches: %w", err)
	}
	return nil
}
