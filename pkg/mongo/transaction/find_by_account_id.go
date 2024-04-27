package transaction

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/mongo"
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

	if err := cursor.All(context.TODO(), &txs); err != nil {
		log.Println("Error decoding transactions: ", err)
		return nil, err
	}

	return txs, nil

}
