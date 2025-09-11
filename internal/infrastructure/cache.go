package infrastructure

import (
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/miztch/valorant-match-schedule/internal/domain"
)

// EventCache is a cache for events
type EventCache struct {
	cache *cache.Cache
}

// NewEventCache creates a new event cache
func NewEventCache() *EventCache {
	eventCache := cache.New(15*time.Minute, cache.NoExpiration)
	return &EventCache{
		cache: eventCache,
	}
}

// GetEventFromCache gets an event from the cache
func (c *EventCache) getEventFromCache(id string) (domain.VlrEvent, bool) {
	if e, found := c.cache.Get(id); found {
		return e.(domain.VlrEvent), true
	}
	return domain.VlrEvent{}, false
}

// CacheEvent caches an event
func (c *EventCache) cacheEvent(e domain.VlrEvent) {
	idStr := strconv.Itoa(e.Id)
	_, exist := c.getEventFromCache(idStr)
	if !exist {
		c.cache.Add(idStr, e, cache.DefaultExpiration)
	}
}
