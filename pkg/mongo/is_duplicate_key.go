package mongo

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsDuplicateKeyError(err error) bool {
	var writeException *mongo.WriteException
	if errors.As(err, &writeException) {
		for _, writeError := range writeException.WriteErrors {
			if writeError.Code == 11000 || writeError.Code == 11001 {
				return true
			}
		}
	}
	return false
}
