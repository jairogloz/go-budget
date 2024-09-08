package core

import (
	"errors"
	"os"
)

type Config struct {
	MongoURI    string
	MongoDBName string
}

// LoadConfig loads the configuration from the environment variables.
//
// Todo: implement better loading of configuration from environment variables.
func LoadConfig() (*Config, error) {
	c := &Config{
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DB_NAME"),
	}
	if c.MongoURI == "" {
		return nil, errors.New("MONGO_URI is required")
	}
	if c.MongoDBName == "" {
		return nil, errors.New("MONGO_DB_NAME is required")
	}
	return c, nil
}
