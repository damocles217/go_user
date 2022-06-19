package app

import (
	"context"
	"encoding/json"

	"github.com/damocles217/user_service/src/user/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUser(id primitive.ObjectID, collection *mongo.Collection) (map[string]any, bool) {
	var user models.User
	var inInterface map[string]interface{}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	err := collection.FindOneAndDelete(context.TODO(), filter).Decode(&user)

	if err != nil {
		return map[string]any{
			"message": "User couldn't be found",
		}, false
	}

	inrec, _ := json.Marshal(&user)
	json.Unmarshal(inrec, &inInterface)

	return inInterface, true
}
