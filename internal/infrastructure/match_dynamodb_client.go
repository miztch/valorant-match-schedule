package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/miztch/valorant-match-schedule/internal/dto"
)

// NewMatchDynamoDBClient creates a DynamoDBClient for Match operations
// using the MATCHLIST_TABLE environment variable
func NewMatchDynamoDBClient(ctx context.Context) (*DynamoDBClient, error) {
	tableName := os.Getenv("MATCHLIST_TABLE")
	if tableName == "" {
		return nil, fmt.Errorf("MATCHLIST_TABLE is not set")
	}
	return NewDynamoDBClient(ctx, tableName)
}

// WriteMatches writes Match items to DynamoDB
func WriteMatches(ctx context.Context, client *DynamoDBClient, matches []dto.Match) error {
	return PutItems(ctx, client, matches)
}
