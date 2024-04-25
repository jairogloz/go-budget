package transaction

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	goBudgetMongo "github.com/jairogloz/go-budget/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository implements the core.TransactionRepository interface using mongo
// as the underlying database.
type repository struct {
	client *mongo.Client
	col    *mongo.Collection
}

// NewRepository creates a new transaction repository.
func NewRepository(client *mongo.Client) core.TransactionRepository {
	return &repository{
		client: client,
		col:    client.Database(goBudgetMongo.DatabaseNameGoBudget).Collection(goBudgetMongo.CollectionNameTransactions),
	}
}
