package transaction

import (
	"context"
	"errors"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Insert inserts a new transaction into the database.
// TODO: this function should be refactor to use a transactional approach: it should only update transaction
func (r repository) Insert(transaction *core.Transaction, newCategory bool) (*core.Account, error) {

	session, err := r.client.StartSession()
	if err != nil {
		log.Println("failed to start session", err.Error())
		return nil, err
	}
	defer session.EndSession(context.Background())

	var updatedAccount core.Account

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		if transaction.ID == nil {
			// Generate new mongo ObjectId
			transaction.ID = primitive.NewObjectID()
		}

		_, err := r.txCol.InsertOne(sessionContext, transaction)
		if err != nil {
			log.Println("failed to insert transaction into database", err.Error())
			return nil, err
		}

		// ==================== UPDATE ACCOUNT BALANCE ====================
		// Create objectId based on account id
		oid, err := primitive.ObjectIDFromHex(transaction.AccountId)
		if err != nil {
			log.Println("failed to create object id", err.Error())
			return nil, err
		}
		filter := bson.M{"_id": oid}
		update := bson.M{"$inc": bson.M{"balance": transaction.Amount}}

		// Use FindOneAndUpdate with ReturnDocument set to After
		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		err = r.accCol.FindOneAndUpdate(sessionContext, filter, update, opts).Decode(&updatedAccount)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				log.Println("account not found", err.Error())
				return nil, err
			}
			log.Println("failed to update account balance", err.Error())
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(context.Background(), callback)
	if err != nil {
		log.Println("transaction failed", err.Error())
		return nil, err
	}

	return &updatedAccount, nil

}
