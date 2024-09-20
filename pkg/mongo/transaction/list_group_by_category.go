package transaction

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// ListGroupByCategory retrieves a list of transactions grouped by category for a user within a specified date range.
// It constructs a MongoDB aggregation pipeline to match transactions based on the user ID and date range, and then groups them by category ID, summing the amounts.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - userID: The ID of the user for whom the transactions are being retrieved.
//   - from: The start date for the transaction list.
//   - to: The end date for the transaction list.
//
// Returns:
//   - []core.TransactionGroupByCategory: The list of transactions grouped by category.
//   - error: An error if there is an issue retrieving the transactions or decoding the results.
func (r repository) ListGroupByCategory(ctx context.Context, userID string, from, to *time.Time) ([]core.TransactionGroupByCategory, error) {
	// MongoDB aggregation pipeline
	pipeline := mongo.Pipeline{
		// Match transactions based on userID and created_at time range
		{
			{"$match", bson.D{
				{"user_id", userID},
				{"created_at", bson.D{
					{"$gte", from},
					{"$lte", to},
				}},
			}},
		},
		// Group by category_id, summing the amounts
		{
			{"$group", bson.D{
				{"_id", bson.D{
					{"$ifNull", bson.A{"$category_id", ""}}, // Group by category or empty string for missing categories
				}},
				{"total_amount", bson.D{{"$sum", "$amount"}}}, // Sum the amounts
			}},
		},
		// Sort by _id
		{
			{"$sort", bson.D{
				{"_id", 1}, // 1 for ascending order, -1 for descending order
			}},
		},
	}

	// Run the aggregation query
	cursor, err := r.txCol.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("error closing cursor:", err)
		}
	}()

	// Parse the results
	var results []struct {
		ID          *string `bson:"_id"`
		TotalAmount float64 `bson:"total_amount"`
	}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	// Convert to the core.TransactionGroupByCategory type
	var groupedTransactions []core.TransactionGroupByCategory
	for _, result := range results {
		groupedTransactions = append(groupedTransactions, core.TransactionGroupByCategory{
			Category: &core.Category{
				ID: *result.ID, // The category ID
			},
			TotalAmount: result.TotalAmount, // The total amount
		})
	}

	return groupedTransactions, nil
}
