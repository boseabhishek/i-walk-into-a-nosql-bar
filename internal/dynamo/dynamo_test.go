package dynamo

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/boseabhishek/i-walk-into-a-nosql-bar/mocks"
	"github.com/stretchr/testify/assert"
)

var tableName string = "table-name"

func initDyanmo() (*Dynamo, *mocks.DynamoDBMock) {
	mockDB := new(mocks.DynamoDBMock)
	dynamo := NewDynamo(mockDB, tableName)

	return dynamo, mockDB
}

func TestGetItem(t *testing.T) {
	ctx := context.TODO()
	input := &dynamodb.GetItemInput{TableName: &tableName}

	t.Run("adds table name to input request", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		inputParam := dynamodb.GetItemInput{}
		mockInput := &dynamodb.GetItemInput{
			TableName: &tableName,
		}
		mockDB.On("GetItemWithContext", ctx, mockInput).Return(&dynamodb.GetItemOutput{}, nil)

		tableAdapter.GetItem(ctx, inputParam)
	})

	t.Run("returns error when there is an aws error", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		mockDB.On("GetItemWithContext", ctx, input).Return(&dynamodb.GetItemOutput{}, errors.New("Failed to query DynamoDB"))

		_, err := tableAdapter.GetItem(ctx, *input)

		assert.Error(t, err)
	})

	t.Run("returns error if item is not found", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		var nilItem map[string]*dynamodb.AttributeValue
		mockDB.On("GetItemWithContext", ctx, input).Return(&dynamodb.GetItemOutput{
			Item: nilItem,
		}, nil)

		_, err := tableAdapter.GetItem(ctx, *input)

		assert.EqualError(t, err, ErrItemNotFound)
	})

	t.Run("successfully returns item and no aws error", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		want := map[string]*dynamodb.AttributeValue{}
		mockDB.On("GetItemWithContext", ctx, input).Return(&dynamodb.GetItemOutput{
			Item: want,
		}, nil)

		got, err := tableAdapter.GetItem(ctx, *input)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}

func TestPutItem(t *testing.T) {
	ctx := context.TODO()
	input := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("123"),
			},
			"testField": {
				S: aws.String("testValue"),
			},
		},
	}

	t.Run("adds table name to input request", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		inputParam := dynamodb.PutItemInput{}
		mockInput := &dynamodb.PutItemInput{
			TableName: &tableName,
		}
		mockDB.On("PutItemWithContext", ctx, mockInput).Return(&dynamodb.PutItemOutput{}, nil)

		tableAdapter.PutItem(ctx, inputParam)
	})

	t.Run("returns error when there is an aws error", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		mockDB.On("PutItemWithContext", ctx, input).Return(&dynamodb.PutItemOutput{}, errors.New("Failed to query DynamoDB"))

		err := tableAdapter.PutItem(ctx, *input)

		assert.Error(t, err)
	})

	t.Run("successfully returns nil if no aws error", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		mockDB.On("PutItemWithContext", ctx, input).Return(&dynamodb.PutItemOutput{}, nil)

		err := tableAdapter.PutItem(ctx, *input)

		assert.Nil(t, err)
	})
}

func TestUpdateItem(t *testing.T) {
	ctx := context.TODO()
	input := &dynamodb.UpdateItemInput{TableName: &tableName}

	t.Run("adds table name to input request", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		inputParam := dynamodb.UpdateItemInput{}
		mockInput := &dynamodb.UpdateItemInput{
			TableName: &tableName,
		}
		mockDB.On("UpdateItemWithContext", ctx, mockInput).Return(&dynamodb.UpdateItemOutput{}, nil)

		tableAdapter.UpdateItem(ctx, inputParam)
	})

	t.Run("returns error when there is an aws error", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		mockDB.On("UpdateItemWithContext", ctx, input).Return(&dynamodb.UpdateItemOutput{}, errors.New("Failed to query DynamoDB"))

		err := tableAdapter.UpdateItem(ctx, *input)

		assert.Error(t, err)
	})

	t.Run("successfully returns nil if no aws error", func(t *testing.T) {
		tableAdapter, mockDB := initDyanmo()
		mockDB.On("UpdateItemWithContext", ctx, input).Return(&dynamodb.UpdateItemOutput{}, nil)

		err := tableAdapter.UpdateItem(ctx, *input)

		assert.Nil(t, err)
	})
}
