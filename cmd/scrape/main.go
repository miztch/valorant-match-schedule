package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miztch/valorant-match-schedule/internal/application"
	"github.com/miztch/valorant-match-schedule/internal/infrastructure"
)

type Payload struct {
	Page int `json:"page"`
}

type Response struct {
	MatchesCount int `json:"matches_count"`
}

// Handler processes match scraping and outputs to SQS
func Handler(ctx context.Context, payload Payload) (string, error) {
	slog.Info("Received scraping request", "page", payload.Page)

	sqsClient, err := infrastructure.NewDefaultSQSClient(ctx)
	if err != nil {
		slog.Error("failed to create SQS client", "error", err.Error())
		return "", err
	}

	matchRepo := infrastructure.NewMatchRepository(infrastructure.NewVlrGGScraper(), sqsClient)
	eventRepo := infrastructure.NewEventRepository(infrastructure.NewVlrGGScraper(), infrastructure.NewEventCache())
	app := application.NewMatchService(matchRepo, eventRepo)

	matches, err := app.FetchMatches(payload.Page)
	if err != nil {
		slog.Error("failed to fetch matches", "error", err.Error())
		return "", err
	}
	slog.Info("Successfully scraped matches", "matchCount", len(matches), "page", payload.Page)

	err = app.WriteMatches(context.Background(), matches)
	if err != nil {
		slog.Error("failed to write matches", "error", err.Error())
		return "", err
	}
	slog.Info("Successfully sent matches to SQS", "matchCount", len(matches))

	response := Response{MatchesCount: len(matches)}
	r, _ := json.Marshal(response)

	return string(r), nil
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	lambda.Start(Handler)
}
