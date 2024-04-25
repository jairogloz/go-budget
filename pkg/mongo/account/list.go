package account

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// List returns the list of accounts for a given user.
func (r repository) List(userId string) ([]core.Account, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOpts := options.Find()
	findOpts.SetLimit(100)
	cursor, err := r.col.Find(ctx, bson.M{"user_id": userId}, findOpts)
	if err != nil {
		fmt.Println("Error finding accounts", err.Error())
		return nil, err
	}

	var accounts []core.Account
	if err = cursor.All(ctx, &accounts); err != nil {
		fmt.Println("Error decoding accounts", err.Error())
		return nil, err
	}

	return accounts, nil
}
