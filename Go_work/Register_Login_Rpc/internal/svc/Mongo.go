package svc

import (
	"Register_Login_Rpc/internal/config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection(c config.Config) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s", c.Mongo.Hosts[0])
	clientOptions := options.Client().ApplyURI(uri)

	clientOptions.SetAuth(options.Credential{
		Username:      c.Mongo.User,
		Password:      c.Mongo.Password,
		AuthSource:    c.Mongo.AuthSource,
		AuthMechanism: "SCRAM-SHA-256",
	})

	clientOptions.SetServerSelectionTimeout(5 * time.Second)
	clientOptions.SetConnectTimeout(5 * time.Second)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("Failed to ping MongoDB: %w", err)
	}

	return client, nil
}
