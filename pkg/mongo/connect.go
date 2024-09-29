package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ConnectMongoDB connects to a MongoDB instance and returns a client to interact
// with the database.
func ConnectMongoDB(mongoURI string) (*mongo.Client, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	disconnect := func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(fmt.Errorf("failed to disconnect from MongoDB: %w", err))
		}
	}

	return client, disconnect, nil
}
