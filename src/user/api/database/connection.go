package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Collection {

	uri_mongo := os.Getenv("URI_MONGO")

	uri := options.Client().ApplyURI(uri_mongo)

	client, err := mongo.Connect(context.TODO(), uri)

	if err != nil {
		println("Error connecting the database: \n", err.Error())
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		println("Error creating the ping: \n", err.Error())
	}

	collection := client.Database("example").Collection("user")

	println("Connected to mongodb succesfully")

	return collection
}
