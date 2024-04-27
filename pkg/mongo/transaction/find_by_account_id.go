package transaction

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/mongo"
	core2 "github.com/jairogloz/go-budget/pkg/mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// FindByAccountID retrieves all the transactions for a given account and userId.
func (r repository) FindByAccountID(userId, accountID string) ([]core.Transaction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), mongo.TimeoutSeconds*time.Second)
	defer cancel()

	// Find all transactions for the given userId and accountId
	// and return them.
	var txs []core.Transaction
	opts := options.Find().SetSort(bson.D{{"created_at", -1}})
	cursor, err := r.txCol.Find(ctx, bson.M{"user_id": userId, "account_id": accountID}, opts)
	if err != nil {
		log.Println("Error finding transactions: ", err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var tx core.Transaction
		if err = cursor.Decode(&tx); err != nil {
			log.Println("Error decoding transaction: ", err)
			return nil, err
		}
		txId, err := core2.ObjectIDToString(tx.ID)
		if err != nil {
			log.Println("Error converting transaction id to string: ", err)
			return nil, err
		}
		tx.ID = txId
		txs = append(txs, tx)
	}

	return txs, nil

}
