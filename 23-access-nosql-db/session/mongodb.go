package session

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// GetMongoClient 调用方需 Disconnect 掉
func GetMongoClient() *mongo.Client {
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//err = client.Connect(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	return client
}
