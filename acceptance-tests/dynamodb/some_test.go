package acceptance_test

import (
	"context"
	"testing"

	dynamo "github.com/boseabhishek/i-walk-into-a-nosql-bar/acceptance-tests/dynamodb/internal"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"
)

const DynamoLocalDbURL = "http://localhost:8000"

func TestTableNoSecondaryIndexExisting(t *testing.T) {

	ctx := context.TODO()

	dynamo.StartOnLocal(t)

	sess := dynamo.CreateSession(DynamoLocalDbURL)

	dynamoClient := dynamodb.New(sess)

	d := dynamo.NewDynamo(dynamoClient, "GameScore")

	input := &dynamodb.CreateTableInput{
		// An array of attributes that describe the key schema for the table and indexes.
		// N.B. Don't include any non-key attribute definitions in AttributeDefinitions.
		// DynamoDB is schemaless (except the key schema)
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UserId"), // Name
				AttributeType: aws.String("S"),      // Type: S|N|B (String|Number|Binary)
			},
			{
				AttributeName: aws.String("GameTitle"),
				AttributeType: aws.String("S"),
			},
			// {
			// 	AttributeName: aws.String("TopScore"),
			// 	AttributeType: aws.String("N"),
			// },
			// {
			// 	AttributeName: aws.String("TopScoreDateTime"),
			// 	AttributeType: aws.String("S"),
			// },
			// {
			// 	AttributeName: aws.String("Wins"),
			// 	AttributeType: aws.String("N"),
			// },
			// {
			// 	AttributeName: aws.String("Losses"),
			// 	AttributeType: aws.String("N"),
			// },
		},
		// Specifies the attributes that make up the primary key for a table or an index.
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UserId"),
				KeyType:       aws.String("HASH"), // HASH - partition key
			},
			{
				AttributeName: aws.String("GameTitle"),
				KeyType:       aws.String("RANGE"), // RANGE - sort key
			},
		},
		// BillingMode as PAY_PER_REQUEST
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		// If you set BillingMode as PROVISIONED, you must specify this property.
		// If you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(d.TableName),
	}
	// GameScore table created.
	d.ManageGameScoreTable(ctx, input, t)

	t.Run("get score by game title and gamer id", func(t *testing.T) {

		gsl := generateGameScores()

		for _, gs := range gsl {
			d.PutItemIntoTable(ctx, gs, t)
		}

		av := map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String("001"),
			},
			"GameTitle": {
				S: aws.String("Call of duty"),
			},
		}

		gi := createGetItemInput(d.TableName, av)

		res := d.GetItemFromTable(ctx, gi, t)

		out := dynamo.GameScore{}

		err := dynamodbattribute.UnmarshalMap(res.Item, &out)
		if err != nil {
			t.Fatalf("setup error: %v", err)
		}

		assert.Equal(t, out.TopScore, 6616)

	})

}

func generateGameScores() []dynamo.GameScore {
	return []dynamo.GameScore{
		{
			UserID:           "001",
			GameTitle:        "Galaxy Invaders",
			TopScore:         5842,
			TopScoreDateTime: "2015-09-15:17:24:38",
			Wins:             10,
			Losses:           12,
		},
		{
			UserID:           "001",
			GameTitle:        "Call of duty",
			TopScore:         6616,
			TopScoreDateTime: "2016-10-15:17:24:38",
			Wins:             9,
			Losses:           1,
		},
		{
			UserID:           "003",
			GameTitle:        "Call of duty",
			TopScore:         3433,
			TopScoreDateTime: "2018-03-20:21:24:38",
			Wins:             10,
			Losses:           5,
		},
	}
}

func createGetItemInput(tableName string, av map[string]*dynamodb.AttributeValue) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       av,
	}
}
