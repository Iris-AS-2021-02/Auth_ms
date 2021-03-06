package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
	uri                      = "mongodb+srv://root:2021@cluster0.34iec.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	uri2                     = "mongodb+srv://root:2021@cluster0.34iec.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

func GetConnection(url string) (*mongo.Client, context.Context, context.CancelFunc) {
	// Create a new client and connect to the server
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}
