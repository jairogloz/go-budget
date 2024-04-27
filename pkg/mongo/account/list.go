package account

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	core2 "github.com/jairogloz/go-budget/pkg/mongo/core"
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
	cursor, err := r.accCol.Find(ctx, bson.M{"user_id": userId}, findOpts)
	if err != nil {
		fmt.Println("Error finding accounts", err.Error())
		return nil, err
	}

	var accounts []core.Account
	// Iterate over the cursor and decode each document
	for cursor.Next(ctx) {
		var account core.Account
		if err = cursor.Decode(&account); err != nil {
			fmt.Println("Error decoding account", err.Error())
			return nil, err
		}
		idString, err := core2.ObjectIDToString(account.ID)
		if err != nil {
			fmt.Println("Error converting account id to string", err.Error())
			return nil, err
		}
		account.ID = idString
		accounts = append(accounts, account)
	}

	return accounts, nil
}
