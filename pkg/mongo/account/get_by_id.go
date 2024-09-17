package account

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// GetByID retrieves an account by its ID for a given user from the MongoDB collection.
// It converts the provided account ID to a MongoDB ObjectID and queries the collection
// for a document matching the account ID and user ID. If found, the account is returned.
//
// Parameters:
//
//	ctx - the context for the request, used for cancellation and deadlines
//	userId - the ID of the user who owns the account
//	id - the ID of the account to retrieve
//
// Returns:
//
//	core.Account - the retrieved account
//	error - an error if one occurred, otherwise nil
func (r repository) GetByID(ctx context.Context, userId, id string) (core.Account, error) {

	// Turn id into an ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return core.Account{}, err
	}

	var account core.Account
	err = r.accCol.FindOne(ctx, bson.M{"_id": oid, "user_id": userId}).Decode(&account)
	if err != nil {
		log.Println(err)
		return core.Account{}, err
	}

	return account, nil
}
