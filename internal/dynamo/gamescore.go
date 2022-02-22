package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GamesScoreDynamo interface lists all the function which can be used to access Dynamodb.
type GamesScoreDynamo interface {
	GetItem(ctx context.Context, i dynamodb.GetItemInput) (map[string]*dynamodb.AttributeValue, error)
}

// GamesScoreRepository implement the GameService here for implementing the necessary methods.
type GamesScoreRepository struct { // TODO: Don't like the name..
	db GamesScoreDynamo
}

// NewGamesScoreRepository returns a new instance of NewGamesScoreRepository
func NewGamesScoreRepository(db GamesScoreDynamo) *GamesScoreRepository {
	return &GamesScoreRepository{db: db}
}

