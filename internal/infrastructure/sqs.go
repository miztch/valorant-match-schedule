package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/miztch/valorant-match-schedule/internal/domain"
)

// SQSClient is a client for SQS
type SQSClient struct {
	client   *sqs.Client
	queueURL string
}

// NewSQSClient creates a new SQSClient
func NewSQSClient(ctx context.Context, queueURL string) (*SQSClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration, %w", err)
	}
	client := sqs.NewFromConfig(cfg)
	return &SQSClient{client: client, queueURL: queueURL}, nil
}

// NewDefaultSQSClient creates a new SQSClient using environment variables
func NewDefaultSQSClient(ctx context.Context) (*SQSClient, error) {
	queueURL := os.Getenv("SQS_QUEUE_URL")
	if queueURL == "" {
		return nil, fmt.Errorf("SQS_QUEUE_URL is not set")
	}
	return NewSQSClient(ctx, queueURL)
}

// WriteMatches writes matches to SQS
func (s *SQSClient) WriteMatches(ctx context.Context, matches []domain.Match) error {
	for _, match := range matches {
		messageBody, err := json.Marshal(match)
		if err != nil {
			return fmt.Errorf("failed to marshal match: %w", err)
		}

		_, err = s.client.SendMessage(ctx, &sqs.SendMessageInput{
			QueueUrl:    &s.queueURL,
			MessageBody: aws.String(string(messageBody)),
		})
		if err != nil {
			return fmt.Errorf("failed to send message to SQS: %w", err)
		}
	}

	return nil
}
