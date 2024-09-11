package transaction

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Delete removes a transaction from the database and updates the associated account balance.
//
// It takes a user ID and a transaction ID as input, starts a MongoDB session, and performs the following steps within a transaction:
//  1. Converts the transaction ID to a MongoDB ObjectID.
//  2. Removes the transaction from the database.
//  3. Updates the account balance associated with the transaction.
//
// If any step fails, it logs the error and returns it.
//
// Parameters:
//   - userId: the ID of the user who owns the transaction
//   - transactionID: the ID of the transaction to be deleted
//
// Returns:
//   - err: an error if there is an issue during the deletion or account balance update
func (r repository) Delete(userId, transactionID string) error {
	session, err := r.client.StartSession()
	if err != nil {
		log.Println("failed to start session", err.Error())
		return err
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		// Convert the transaction ID to an object ID.
		objID, err := primitive.ObjectIDFromHex(transactionID)
		if err != nil {
			log.Println("error converting transaction ID to object ID:", err)
			return nil, err
		}

		// Remove the transaction from the database.
		transaction := core.Transaction{}
		err = r.txCol.FindOneAndDelete(sessionContext, bson.M{"_id": objID, "user_id": userId}).Decode(&transaction)
		if err != nil {
			log.Println("error deleting transaction:", err)
			return nil, err
		}

		// Update the account balance.
		oid, err := primitive.ObjectIDFromHex(*transaction.AccountId)
		if err != nil {
			log.Println("failed to create object id", err.Error())
			return nil, err
		}
		filter := bson.M{"_id": oid}
		update := bson.M{"$inc": bson.M{"balance": -transaction.Amount}}
		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		var updatedAccount core.Account
		err = r.accCol.FindOneAndUpdate(sessionContext, filter, update, opts).Decode(&updatedAccount)
		if err != nil {
			log.Println("failed to update account balance", err.Error())
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(context.Background(), callback)
	if err != nil {
		log.Println("transaction failed", err.Error())
		return err
	}

	return nil
}
