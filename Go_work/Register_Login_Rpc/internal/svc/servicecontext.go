package svc

import (
	"context"
	"fmt"
	"time"

	"Register_Login_Rpc/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Config  config.Config
	MongoDB *mongo.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := connection(c)
	if err != nil {
		panic(fmt.Sprintf("Failed to create MongoDB connection: %v", err))
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to ping MongoDB: %v", err))
	}

	return &ServiceContext{
		Config:  c,
		MongoDB: client,
	}
}
