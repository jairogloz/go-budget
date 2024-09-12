package transaction

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// List retrieves a list of transactions for a user within a specified date range and with pagination options.
// It constructs a MongoDB query filter based on the user ID and date range, and applies pagination and sorting options.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - userID: The ID of the user for whom the transactions are being retrieved.
//   - from: The start date for the transaction list.
//   - to: The end date for the transaction list.
//   - limit: The maximum number of transactions to retrieve.
//   - offset: The number of transactions to skip for pagination.
//
// Returns:
//   - core.Transactions: The list of transactions.
//   - error: An error if there is an issue retrieving the transactions or decoding the results.
func (r repository) List(ctx context.Context, userID string, from, to *time.Time, limit, offset int) (core.Transactions, error) {

	filter := bson.M{
		"user_id": userID,
		"created_at": bson.M{
			"$gte": from,
			"$lte": to,
		},
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	findOptions.SetSort(bson.D{{"created_at", -1}})

	cursor, err := r.txCol.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("error closing cursor:", err)
		}
	}()

	var transactions core.Transactions
	for cursor.Next(ctx) {
		var transaction core.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return transactions, nil

}
