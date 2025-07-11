package infrastructure

import (
	"fmt"
	"strings"

	"github.com/miztch/valorant-match-schedule/internal/domain"
)

// EventRepository is an interface for getting events
type eventRepositoryImpl struct {
	vlrGGScraper *VlrGGScraper
	eventCache   *EventCache
}

// NewEventRepository creates a new event repository
func NewEventRepository(vlrGGScraper *VlrGGScraper, eventCache *EventCache) *eventRepositoryImpl {
	return &eventRepositoryImpl{
		vlrGGScraper: vlrGGScraper,
		eventCache:   eventCache,
	}
}

// GetEvent gets an event
func (r *eventRepositoryImpl) GetEvent(eventPagePath string) (domain.VlrEvent, error) {
	id := strings.Split(eventPagePath, "/")[2]

	event, found := r.eventCache.getEventFromCache(id)
	if found {
		return event, nil
	}

	event, err := r.vlrGGScraper.scrapeEvent(eventPagePath)
	if err != nil {
		return domain.VlrEvent{}, fmt.Errorf("failed to get event: %w", err)
	}
	r.eventCache.cacheEvent(event)

	return event, nil
}
