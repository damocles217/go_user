package app

import (
	"context"
	"encoding/json"

	"github.com/damocles217/user_service/src/user/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(filterId any, filterType string, collection *mongo.Collection, opts *options.FindOneOptions) (map[string]any, bool) {
	var user models.User
	// Translate the struct values to map
	var inInterface map[string]interface{}

	if filterType == "" {
		// Default value to filter
		filterType = "_id"
	}

	filter := bson.D{primitive.E{Key: filterType, Value: filterId}}

	err := collection.FindOne(context.TODO(), filter, opts).Decode(&user)

	if err != nil {
		return map[string]interface{}{
			"message": "userNotFound",
		}, false
	}

	// Translate for struct to map
	inrec, _ := json.Marshal(user)
	json.Unmarshal(inrec, &inInterface)

	return inInterface, true
}
