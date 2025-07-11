package domain

import (
	"context"

	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// MatchWriter defines the interface for writing match data
type MatchWriter interface {
	// WriteMatches writes multiple matches to the data store
	WriteMatches(ctx context.Context, matches []dto.Match) error
}
