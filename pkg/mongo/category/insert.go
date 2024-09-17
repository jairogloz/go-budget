package category

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	core2 "github.com/jairogloz/go-budget/pkg/mongo/category/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// Insert inserts a new category into the database.
func (r repository) Insert(ctx context.Context, category *core.Category) (*core.Category, error) {
	if category == nil {
		return nil, fmt.Errorf("category is nil")
	}

	objectID := primitive.NewObjectID()
	mongoCategory := core2.MongoCategory{
		ID:       objectID,
		Category: *category,
	}

	_, err := r.catCol.InsertOne(ctx, mongoCategory)
	if err != nil {
		// If error is due to duplicate key, return nil and error
		if mongo.IsDuplicateKeyError(err) {
			log.Println("category already exists", err.Error())
			return nil, nil
		}
		return nil, err
	}
	category.ID = objectID.Hex()

	return category, nil
}
