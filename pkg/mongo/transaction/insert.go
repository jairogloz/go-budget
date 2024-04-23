package transaction

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// Insert inserts a new transaction into the database.
func (r repository) Insert(transaction *core.Transaction, newCategory bool) error {

	if transaction.ID == "" {
		// Generate new mongo ObjectId
		transaction.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.col.InsertOne(context.Background(), transaction)
	if err != nil {
		log.Println("failed to insert transaction into database", err.Error())
		return err
	}

	return nil

}
