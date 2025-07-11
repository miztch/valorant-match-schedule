package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miztch/valorant-match-schedule/internal/application"
	"github.com/miztch/valorant-match-schedule/internal/infrastructure"
)

// Handler processes SQS events containing match data
func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	slog.Info("Received SQS event", "recordCount", len(sqsEvent.Records))

	// Log each record for debugging
	for i, record := range sqsEvent.Records {
		slog.Info("Processing SQS record", "index", i, "messageId", record.MessageId, "source", record.EventSource)
	}

	// Create MatchWriter
	matchWriter, err := infrastructure.NewMatchWriter(ctx)
	if err != nil {
		slog.Error("Failed to create match writer", "error", err.Error())
		return err
	}

	// Create MatchRegistrationService
	service := application.NewMatchRegistrationService(matchWriter)

	// Process SQS event
	slog.Info("Starting SQS event processing")
	err = service.ProcessSQSEvent(ctx, sqsEvent)
	if err != nil {
		slog.Error("Failed to process SQS event", "error", err.Error())
		return err
	}
	slog.Info("Successfully completed SQS event processing", "recordCount", len(sqsEvent.Records))
	return nil
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	lambda.Start(Handler)
}
