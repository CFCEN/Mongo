package Db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient Connect uri is mongodb's uri , mongodb://localhost:27017
func CreateClient(uri string) *mongo.Client {
	var ctx = context.TODO()
	url := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, url)
	fmt.Println("Connected to MongoDB!")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
