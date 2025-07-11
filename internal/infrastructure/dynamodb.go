package infrastructure

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDB BatchWriteItem has a limit of 25 items
const batchSize = 25

// DynamoDBClient is a generic client for DynamoDB operations
type DynamoDBClient struct {
	client    *dynamodb.Client
	TableName string
}

// NewDynamoDBClient creates a new DynamoDBClient
func NewDynamoDBClient(ctx context.Context, tableName string) (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration, %w", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{client: client, TableName: tableName}, nil
}

// BatchWrite writes items to DynamoDB using BatchWriteItem
func (d *DynamoDBClient) BatchWrite(ctx context.Context, writeReqs []types.WriteRequest) error {
	_, err := d.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{d.TableName: writeReqs},
	})
	if err != nil {
		return fmt.Errorf("failed to write batch: %w", err)
	}
	return nil
}

// PutItems writes items to DynamoDB in batches
// T can be any struct that can be marshaled to DynamoDB attributes
func PutItems[T any](ctx context.Context, client *DynamoDBClient, items []T) error {
	var allWriteReqs []types.WriteRequest

	// Prepare all write requests
	for _, item := range items {
		marshaledItem, err := attributevalue.MarshalMap(item)
		if err != nil {
			return fmt.Errorf("failed to marshal item: %w", err)
		}
		allWriteReqs = append(allWriteReqs, types.WriteRequest{
			PutRequest: &types.PutRequest{Item: marshaledItem},
		})
	}

	// Write in batches
	for i := 0; i < len(allWriteReqs); i += batchSize {
		end := i + batchSize
		if end > len(allWriteReqs) {
			end = len(allWriteReqs)
		}

		chunk := allWriteReqs[i:end]
		err := client.BatchWrite(ctx, chunk)
		if err != nil {
			return fmt.Errorf("failed to write batch: %w", err)
		}
	}

	return nil
}

// GetItem retrieves a single item from DynamoDB
func (d *DynamoDBClient) GetItem(ctx context.Context, key map[string]types.AttributeValue) (*dynamodb.GetItemOutput, error) {
	input := &dynamodb.GetItemInput{
		TableName: &d.TableName,
		Key:       key,
	}
	return d.client.GetItem(ctx, input)
}

// QueryItems performs a query operation on DynamoDB
func (d *DynamoDBClient) QueryItems(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	input.TableName = &d.TableName
	return d.client.Query(ctx, input)
}

// ScanItems performs a scan operation on DynamoDB
func (d *DynamoDBClient) ScanItems(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	input.TableName = &d.TableName
	return d.client.Scan(ctx, input)
}

// GetClient returns the underlying DynamoDB client for advanced operations
func (d *DynamoDBClient) GetClient() *dynamodb.Client {
	return d.client
}

// GetTableName returns the table name
func (d *DynamoDBClient) GetTableName() string {
	return d.TableName
}
