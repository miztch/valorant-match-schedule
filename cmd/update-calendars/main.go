package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miztch/valorant-match-schedule/internal/application"
)

// Lambda function that receives DynamoDB Stream events and creates Google Calendar events
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	slog.Info("Received DynamoDB Stream event", "recordCount", len(event.Records))

	// Log each record for debugging
	for i, record := range event.Records {
		slog.Info("Processing DynamoDB Stream record",
			"index", i,
			"eventName", record.EventName,
			"eventSource", record.EventSource,
			"eventSourceARN", record.EventSourceArn,
		)
	}

	calendarService, err := application.NewCalendarServiceFromEnv(ctx)
	if err != nil {
		slog.Error("Failed to create calendar service", "error", err.Error())
		return err
	}

	err = calendarService.HandleDynamoDBStream(ctx, event)
	if err != nil {
		slog.Error("Failed to process DynamoDB Stream event", "error", err.Error())
		return err
	}
	return nil
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	lambda.Start(Handler)
}
