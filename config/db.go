package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConn() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	dsn := os.Getenv("DSN")

	fmt.Println("dsn:", dsn)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
