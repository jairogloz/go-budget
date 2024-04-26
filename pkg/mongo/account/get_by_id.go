package account

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func (r repository) GetByID(userId, id string) (core.Account, error) {

	// Turn id into an ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return core.Account{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var account core.Account
	err = r.col.FindOne(ctx, bson.M{"_id": oid, "user_id": userId}).Decode(&account)
	if err != nil {
		log.Println(err)
		return core.Account{}, err
	}

	return account, nil
}
