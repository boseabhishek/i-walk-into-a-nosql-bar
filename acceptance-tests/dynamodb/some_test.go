package acceptance_test

import (
	"context"
	"log"
	"sort"
	"testing"

	dynamo "github.com/boseabhishek/i-walk-into-a-nosql-bar/acceptance-tests/dynamodb/internal"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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
				AttributeType: aws.String("S"),      // Type: S|N|B (String|Number|Binary) Also, Map?
			},
			{
				AttributeName: aws.String("GameTitle"),
				AttributeType: aws.String("S"),
			},
		},
		// Specifies the attributes that make up the primary key for a table or an index.
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UserId"),
				KeyType:       aws.String("HASH"), // HASH - partition key
			},
			// partition key is very important for horizontal scaling
			{
				AttributeName: aws.String("GameTitle"),
				KeyType:       aws.String("RANGE"), // RANGE - sort key
			},
			// This makes a composite primary key of partition key + sort key
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

	gsl := generateGameScores()

	// Adds data to table.
	for _, gs := range gsl {
		d.PutItemIntoTable(ctx, gs, t)
	}

	t.Run("get score by game title and gamer id", func(t *testing.T) {

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
			t.Fatalf("error unmarshalling getItem: %v", err)
		}

		assert.Equal(t, out.TopScore, 6616)

	})
	t.Run("get topscores by game title", func(t *testing.T) {

		// Create the Expression to fill the input struct with.
		// Get all GameScores for that game title
		filt := expression.Name("GameTitle").Equal(expression.Value("Call of duty"))

		// Get back the TopScore
		proj := expression.NamesList(expression.Name("TopScore"))

		si := creatScanInput(d.TableName, filt, proj)

		res := d.ScanItemFromTable(ctx, si, t)

		var outs []dynamo.GameScore

		for _, i := range res.Items {
			out := dynamo.GameScore{}

			err := dynamodbattribute.UnmarshalMap(i, &out)

			if err != nil {
				t.Fatalf("error unmarshalling scan: %v", err)
			}

			outs = append(outs, out)
		}

		assert.Equal(t, len(outs), 3)

	})
	t.Run("find gamer ids by highest score for a game", func(t *testing.T) {

		// Create the Expression to fill the input struct with.
		// Get all GameScores for that game title

		forGame := "Call of duty"

		filt := expression.Name("GameTitle").Equal(expression.Value(forGame))

		// Get back the TopScore
		proj := expression.NamesList(expression.Name("UserId"), expression.Name("TopScore"))

		si := creatScanInput(d.TableName, filt, proj)

		res := d.ScanItemFromTable(ctx, si, t)

		var outs []dynamo.GameScore

		for _, i := range res.Items {
			out := dynamo.GameScore{}

			err := dynamodbattribute.UnmarshalMap(i, &out)

			if err != nil {
				t.Fatalf("error unmarshalling scan: %v", err)
			}
			outs = append(outs, out)
		}

		assert.Equal(t, len(outs), 3)

		// Sort outs by descending topscore.
		sort.SliceStable(outs, func(i, j int) bool {
			return outs[i].TopScore > outs[j].TopScore
		})

		assert.Equal(t, outs[0].UserID, "001")

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
		{
			UserID:           "002",
			GameTitle:        "Call of duty",
			TopScore:         4211,
			TopScoreDateTime: "2017-03-20:21:24:38",
			Wins:             14,
			Losses:           8,
		},
	}
}

func createGetItemInput(tableName string, av map[string]*dynamodb.AttributeValue) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       av,
	}
}

func creatScanInput(tableName string, filt expression.ConditionBuilder, proj expression.ProjectionBuilder) *dynamodb.ScanInput {

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	// Build the query input parameters
	return &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}
}
