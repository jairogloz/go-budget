package core

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoTransaction is a helper type that wraps domain/core/Transaction but
// overrides the id field in order to be able to set a mongo ObjectID.
type MongoTransaction struct {
	ID               primitive.ObjectID `bson:"_id"`
	core.Transaction `bson:",inline"`
}
