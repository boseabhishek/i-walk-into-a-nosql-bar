package dynamo

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// ErrItemNotFound ...
const ErrItemNotFound = "item not found in dynamo"

// Dynamo is a DynamoDB connection to a specific table.
type Dynamo struct {
	dynamo    dynamodbiface.DynamoDBAPI
	tableName string
}

// NewDynamo returns a new instance of Dynamo.
func NewDynamo(db dynamodbiface.DynamoDBAPI, t string) *Dynamo {
	return &Dynamo{
		dynamo:    db,
		tableName: t,
	}
}

// GetItem gets an item from the DynamoDB table.
func (r *Dynamo) GetItem(ctx context.Context, i dynamodb.GetItemInput) (map[string]*dynamodb.AttributeValue, error) {
	i.TableName = &r.tableName
	resp, err := r.dynamo.GetItemWithContext(ctx, &i)

	if err != nil {
		return nil, err
	}

	if resp.Item == nil {
		return nil, errors.New(ErrItemNotFound)
	}

	return resp.Item, nil
}

// PutItem puts an item in the DynamoDB table
func (r *Dynamo) PutItem(ctx context.Context, i dynamodb.PutItemInput) error {
	i.TableName = &r.tableName
	_, err := r.dynamo.PutItemWithContext(ctx, &i)

	if err != nil {
		return err
	}

	return nil
}

// UpdateItem updates an item in the DynamoDB table
func (r *Dynamo) UpdateItem(ctx context.Context, i dynamodb.UpdateItemInput) error {
	i.TableName = &r.tableName
	_, err := r.dynamo.UpdateItemWithContext(ctx, &i)

	if err != nil {
		return err
	}

	return nil
}
