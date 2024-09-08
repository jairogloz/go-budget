package account

import (
	"context"
	"fmt"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"github.com/jairogloz/go-budget/pkg/mongo/account/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Create creates a new account.
func (r repository) Create(account domainCore.Account) (insertedID string, err error) {
	
	objectID := primitive.NewObjectID()

	mongoAccount := core.MongoAccount{
		Account: account,
		ID:      objectID,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), mongo.TimeoutSeconds*time.Second)
	defer cancel()

	log.Println("about to insert account into database")

	insertOneResult, err := r.accCol.InsertOne(ctx, mongoAccount)
	if err != nil {
		log.Println("failed to insert account into database", err.Error())
		return "", err
	}

	id, ok := insertOneResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("failed to convert inserted ID to object ID")
		return "", fmt.Errorf("failed to convert inserted ID to object ID")
	}

	return id.Hex(), nil
}
