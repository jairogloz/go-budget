package account

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Create creates a new account.
func (r repository) Create(account *core.Account) error {

	if account.ID == nil {
		// Generate new mongo ObjectId
		account.ID = primitive.NewObjectID()
	}

	ctx, cancel := context.WithTimeout(context.TODO(), mongo.TimeoutSeconds*time.Second)
	defer cancel()

	log.Println("about to insert account into database")

	_, err := r.col.InsertOne(ctx, account)
	if err != nil {
		log.Println("failed to insert account into database", err.Error())
	}

	return err
}
