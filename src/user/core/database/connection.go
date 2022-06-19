package database

import (
	"context"
	"os"

	"github.com/damocles217/user_service/src/user/core/validators"
	"go.mongodb.org/mongo-driver/bson"
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

	// validators
	validator := validators.UserValidator()
	opts := options.CreateCollection().SetValidator(validator)

	err = client.Database("example").CreateCollection(context.TODO(), "user", opts)

	if err != nil {
		println("Collection created:\n", err.Error())
	}
	collection := client.Database("example").Collection("user")

	// * Setting index for unique values
	// index in ascending order or -1 for descending order

	uniqueValuesIndex := [3]string{"email", "code_auth", "userId"}
	for _, value := range uniqueValuesIndex {
		uniqueValue := mongo.IndexModel{
			Keys:    bson.M{value: 1},
			Options: options.Index().SetUnique(true),
		}

		_, err = collection.Indexes().CreateOne(context.TODO(), uniqueValue)
	}

	if err != nil {
		println("No created index:\n", err.Error())
	}
	println("Connected to mongodb succesfully")

	return collection
}
