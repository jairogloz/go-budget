package account

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
	goBudgetMongo "github.com/jairogloz/go-budget/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository implements the core.AccountRepository interface using mongo
// as the underlying database.
type repository struct {
	client *mongo.Client
	accCol *mongo.Collection
	txCol  *mongo.Collection
}

// NewRepository creates a new account repository.
func NewRepository(client *mongo.Client, c *core.Config) ports.AccountRepository {
	return &repository{
		client: client,
		accCol: client.Database(c.MongoDBName).Collection(goBudgetMongo.CollectionNameAccounts),
		txCol:  client.Database(c.MongoDBName).Collection(goBudgetMongo.CollectionNameTransactions),
	}
}
