package core

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDToString converts an object id to a string.
func ObjectIDToString(id interface{}) (string, error) {
	// Cast the id to a primitive.ObjectID
	// and return the string representation of it.
	if objectID, ok := id.(primitive.ObjectID); ok {
		return objectID.Hex(), nil
	}

	return "", fmt.Errorf("failed to cast id to primitive.ObjectID, got %T", id)
}
