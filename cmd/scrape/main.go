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

type MatchesSummary struct {
	Total     int `json:"total"`
	Succeeded int `json:"succeeded"`
	Failed    int `json:"failed"`
}

type Response struct {
	Matches MatchesSummary `json:"matches"`
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

	matches, matchURLsCount, err := app.FetchMatches(payload.Page)
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

	response := Response{
		Matches: MatchesSummary{
			Total:     matchURLsCount,
			Succeeded: len(matches),
			Failed:    matchURLsCount - len(matches),
		},
	}
	r, _ := json.Marshal(response)

	return string(r), nil
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	lambda.Start(Handler)
}
