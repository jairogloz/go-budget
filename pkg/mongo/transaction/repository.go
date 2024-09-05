package transaction

import (
	"github.com/jairogloz/go-budget/pkg/domain/ports"
	goBudgetMongo "github.com/jairogloz/go-budget/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository implements the core.TransactionRepository interface using mongo
// as the underlying database.
type repository struct {
	accCol *mongo.Collection
	client *mongo.Client
	txCol  *mongo.Collection
}

// NewRepository creates a new transaction repository.
func NewRepository(client *mongo.Client) ports.TransactionRepository {
	return &repository{
		accCol: client.Database(goBudgetMongo.DatabaseNameGoBudget).Collection(goBudgetMongo.CollectionNameAccounts),
		client: client,
		txCol:  client.Database(goBudgetMongo.DatabaseNameGoBudget).Collection(goBudgetMongo.CollectionNameTransactions),
	}
}
