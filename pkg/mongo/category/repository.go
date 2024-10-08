package category

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
	goBudgetMongo "github.com/jairogloz/go-budget/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository implements the core.CategoryRepository interface using mongo
// as the underlying database.
type repository struct {
	client *mongo.Client
	catCol *mongo.Collection
}

// NewRepository creates a new category repository.
func NewRepository(client *mongo.Client, c *core.Config) ports.CategoryRepository {
	return &repository{
		client: client,
		catCol: client.Database(c.MongoDBName).Collection(goBudgetMongo.CollectionNameCategories),
	}
}
