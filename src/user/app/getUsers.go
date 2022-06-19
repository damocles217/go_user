package app

import (
	"context"

	"github.com/damocles217/user_service/src/user/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(collection *mongo.Collection, pageInt int64) ([]models.User, bool) {
	var users []models.User

	filter := bson.D{}

	option := bson.D{
		{"name", 1},
		{"lastname", 1},
		{"userId", 1},
	}

	opts := options.Find().SetProjection(option).SetLimit(3).SetSkip(3 * pageInt)
	cursor, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		return users, false
	}

	for cursor.Next(context.TODO()) {
		var user models.User
		err = cursor.Decode(&user)
		if err != nil {
			return users, false
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return users, false
	}

	cursor.Close(context.TODO())

	return users, true
}
