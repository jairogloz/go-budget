package account

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Delete deletes an account.
func (r repository) Delete(userId, id string) error {

	// Convert the id string to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Create the filter to delete the account
	filter := bson.M{"_id": objectID, "user_id": userId}

	ctx, cancel := context.WithTimeout(context.TODO(), mongo.TimeoutSeconds*time.Second)
	defer cancel()

	log.Println("about to delete account from database")

	_, err = r.col.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("failed to delete account from database", err.Error())
	}

	return err
}
