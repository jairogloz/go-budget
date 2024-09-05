package transaction

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	core2 "github.com/jairogloz/go-budget/pkg/mongo/transaction/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Insert inserts a new transaction into the database.
// It takes a pointer to a core.Transaction as input and returns an error if the insertion fails.
//
// If the transaction is nil, it returns an error indicating that the transaction is nil.
// It creates a context with a timeout of 15 seconds for the database operation.
// It generates a new MongoDB ObjectID for the transaction.
// It wraps the core.Transaction in a MongoTransaction struct with the generated ObjectID.
// It attempts to insert the MongoTransaction into the database collection.
// If the insertion fails, it logs the error and returns it.
// If the insertion is successful, it returns nil.
func (r repository) Insert(transaction *core.Transaction) error {
	if transaction == nil {
		return fmt.Errorf("transaction is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	txObjectID := primitive.NewObjectID()

	mongoTx := core2.MongoTransaction{
		ID:          txObjectID,
		Transaction: *transaction,
	}

	_, err := r.txCol.InsertOne(ctx, mongoTx)
	if err != nil {
		log.Println("failed to insert transaction into database", err.Error())
		return err
	}

	return nil
}
