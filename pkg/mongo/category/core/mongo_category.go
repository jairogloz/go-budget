package core

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoCategory is a helper type that wraps domain/core/Category but
// overrides the id field in order to be able to set a mongo ObjectID.
type MongoCategory struct {
	ID            primitive.ObjectID `bson:"_id"`
	core.Category `bson:",inline"`
}
