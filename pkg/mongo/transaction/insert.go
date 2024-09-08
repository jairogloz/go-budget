package transaction

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	transactionCore "github.com/jairogloz/go-budget/pkg/mongo/transaction/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Insert inserts a new transaction into the database.
//
// It takes a pointer to a core.Transaction as input and returns the inserted ID as a string and an error if the insertion fails.
//
// If the transaction is nil, it returns an error indicating that the transaction is nil.
// It creates a new MongoDB ObjectID for the transaction and wraps the core.Transaction in a MongoTransaction struct with the generated ObjectID.
// It starts a MongoDB session and attempts to insert the MongoTransaction into the database collection within a transaction.
// If the insertion fails, it logs the error and returns it.
// If the insertion is successful, it updates the account balance associated with the transaction.
// If the account balance update fails, it logs the error and returns it.
// If the account balance update is successful, it commits the transaction and returns the inserted ID as a string.
// If the inserted ID cannot be converted to a MongoDB ObjectID, it logs the error and returns it.
//
// Parameters:
//   - transaction: a pointer to the core.Transaction to be inserted
//
// Returns:
//   - insertedID: the ID of the inserted transaction as a string
//   - err: an error if there is an issue during insertion or account balance update
func (r repository) Insert(transaction *core.Transaction) (insertedID string, err error) {
	if transaction == nil {
		return "", fmt.Errorf("transaction is nil")
	}

	txObjectID := primitive.NewObjectID()

	mongoTx := transactionCore.MongoTransaction{
		ID:          txObjectID,
		Transaction: *transaction,
	}

	session, err := r.client.StartSession()
	if err != nil {
		log.Println("failed to start session", err.Error())
		return "", fmt.Errorf("failed to start session: %w", err)
	}
	defer session.EndSession(context.Background())

	var insertOneResult *mongo.InsertOneResult
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		insertOneResult, err = r.txCol.InsertOne(sessionContext, mongoTx)
		if err != nil {
			log.Println("failed to insert transaction into database", err.Error())
			return insertedID, err
		}

		// Update the account balance.
		oid, err := primitive.ObjectIDFromHex(*transaction.AccountId)
		if err != nil {
			log.Println("failed to create object id", err.Error())
			return nil, err
		}
		filter := bson.M{"_id": oid}
		update := bson.M{"$inc": bson.M{"balance": transaction.Amount}}
		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		var updatedAccount core.Account
		err = r.accCol.FindOneAndUpdate(sessionContext, filter, update, opts).Decode(&updatedAccount)
		if err != nil {
			log.Println("failed to update account balance", err.Error())
			return nil, err
		}

		return nil, nil

	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		log.Println("failed to execute transaction", err.Error())
		return "", fmt.Errorf("failed to execute transaction: %w", err)
	}

	id, ok := insertOneResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("failed to convert inserted ID to object ID")
		return "", fmt.Errorf("failed to convert inserted ID to object ID")
	}

	return id.Hex(), nil
}
