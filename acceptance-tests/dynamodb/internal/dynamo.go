package dynamo

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
)

// Dynamo ...
type Dynamo struct {
	dynamodbiface.DynamoDBAPI

	TableName string
}

// NewDynamo ...
func NewDynamo(sess dynamodbiface.DynamoDBAPI, tableName string) *Dynamo {
	return &Dynamo{
		DynamoDBAPI: sess,
		TableName:   tableName,
	}
}

// StartOnLocal ...
func StartOnLocal(t *testing.T) {

	// Helper marks the calling function as a test helper function.
	// When printing file and line information, that function will be skipped.
	// Helper may be called simultaneously from multiple goroutines.
	t.Helper()

	// docker-compose -f docker-compose.yml up -d
	compose := testcontainers.NewLocalDockerCompose(
		[]string{"docker-compose.yml"},
		strings.ToLower(uuid.New().String()),
	)

	invokeErr := compose.WithCommand([]string{"up", "-d"}).Invoke()

	if invokeErr.Error != nil {
		t.Fatalf("error invoking dokcer compose: %v", invokeErr.Error)
	}

	// T.Cleanup() as an "improved" and extended version of defer.
	// It also documents that the passed functions are for cleanup purposes.
	t.Cleanup(func() {
		compose.Down()
	})
}

// CreateSession creates an AWS session based on Localstack configs.
func CreateSession(url string) (sess *session.Session) {

	cfg := aws.NewConfig().
		WithRegion("eu-west-1").
		WithEndpoint(url).
		WithCredentials(credentials.NewStaticCredentials("test", "test", ""))
	sess = session.Must(session.NewSession(cfg))

	return
}

// ManageGameScoreTable ...
func (d *Dynamo) ManageGameScoreTable(ctx context.Context, input *dynamodb.CreateTableInput, t *testing.T) {

	t.Helper()

	_, err := d.CreateTableWithContext(ctx, input)
	if err != nil {
		t.Fatalf("error creating dynamo table: %v", err)
	}

	t.Cleanup(func() {
		_, err := d.DeleteTableWithContext(ctx, &dynamodb.DeleteTableInput{
			TableName: &d.TableName,
		})
		if err != nil {
			t.Fatalf("error deleting dynamo table: %v", err)
		}
	})

}

// PutItemIntoTable ...
func (d *Dynamo) PutItemIntoTable(ctx context.Context, v interface{}, t *testing.T) {

	t.Helper()

	av, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		t.Fatalf("error marshalling data in PutItemIntoTable: %v", err)
	}

	// Create item in table Movies
	pi := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.TableName),
	}

	_, err = d.PutItem(pi)
	if err != nil {
		t.Fatalf("error putting item into dynamo table: %v", err)
	}
}

// GetItemFromTable ...
func (d *Dynamo) GetItemFromTable(ctx context.Context, input *dynamodb.GetItemInput, t *testing.T) (res *dynamodb.GetItemOutput) {

	t.Helper()

	res, err := d.GetItem(input)
	if err != nil {
		t.Fatalf("error getitem from table: %v", err)
	}

	return

}

// ScanItemFromTable ...
func (d *Dynamo) ScanItemFromTable(ctx context.Context, params *dynamodb.ScanInput, t *testing.T) (res *dynamodb.ScanOutput) {

	t.Helper()

	res, err := d.Scan(params)
	if err != nil {
		log.Fatalf("error scan from table: %v", err)
	}

	return

}
