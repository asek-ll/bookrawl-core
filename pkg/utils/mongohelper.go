package utils

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateMongoClientWithUri(connUri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connUri))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateMongoClient() (*mongo.Client, error) {
	connUri := os.Getenv("MONGODB_URI")

	if connUri == "" {
		return nil, errors.New("mongodbURI required")
	}

	return CreateMongoClientWithUri(connUri)
}
