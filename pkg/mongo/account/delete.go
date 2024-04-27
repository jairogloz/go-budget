package account

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// Delete deletes an account.
func (r repository) Delete(userId, id string) error {
	session, err := r.client.StartSession()
	if err != nil {
		log.Println("failed to start session", err.Error())
		return err
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		// Convert the id string to a primitive.ObjectID
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		// Create the filter to delete the account
		filter := bson.M{"_id": objectID, "user_id": userId}

		log.Println("about to delete account from database")

		_, err = r.accCol.DeleteOne(sessionContext, filter)
		if err != nil {
			log.Println("failed to delete account from database", err.Error())
			return nil, err
		}

		// Delete all transactions associated with the account
		_, err = r.txCol.DeleteMany(sessionContext, bson.M{"account_id": id})
		if err != nil {
			log.Println("failed to delete transactions associated with the account", err.Error())
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
