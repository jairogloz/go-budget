package core

import (
	"errors"
	"os"
)

type Config struct {
	MongoURI string
}

// LoadConfig loads the configuration from the environment variables.
//
// Todo: implement better loading of configuration from environment variables.
func LoadConfig() (*Config, error) {
	c := &Config{
		MongoURI: os.Getenv("MONGO_URI"),
	}
	if c.MongoURI == "" {
		return nil, errors.New("MONGO_URI is required")
	}
	return c, nil
}
