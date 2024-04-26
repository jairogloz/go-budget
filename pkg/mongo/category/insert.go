package category

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// Insert inserts a new category into the database.
func (r repository) Insert(category *core.Category) (*core.Category, error) {

	if category.ID == nil {
		category.ID = primitive.NewObjectID()
	}

	_, err := r.catCol.InsertOne(context.TODO(), category)
	if err != nil {
		// If error is due to duplicate key, return nil and error
		if mongo.IsDuplicateKeyError(err) {
			log.Println("category already exists", err.Error())
			return nil, nil
		}
		return nil, err
	}

	return category, nil
}
