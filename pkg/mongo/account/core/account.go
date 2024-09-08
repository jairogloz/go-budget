package core

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoAccount is a helper type that wraps a core.Account and overrides the
// id (string) field with the ObjectID.
type MongoAccount struct {
	core.Account `bson:",inline"`
	ID           primitive.ObjectID `bson:"_id"`
}
