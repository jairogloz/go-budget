package account

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

// CountAccounts returns the number of accounts for a given user.
func (r repository) CountAccounts(userId string) (int, error) {
	colCount, err := r.accCol.CountDocuments(context.Background(), bson.M{"user_id": userId})
	if err != nil {
		return 0, fmt.Errorf("failed to count accounts: %w", err)
	}

	return int(colCount), nil
}
