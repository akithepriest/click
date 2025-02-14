package database

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context) (*mongo.Client, error){
	mongoDBConnURI := os.Getenv("MONGODB_CONN_STRING")

	if mongoDBConnURI == "" {
		return nil, errors.New("MONGODB_CONN_STRING has not been defined")
	}
	clientOptions := options.Client().ApplyURI(mongoDBConnURI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}