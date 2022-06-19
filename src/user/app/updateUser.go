package app

import (
	"context"
	"encoding/json"

	"github.com/damocles217/user_service/src/user/api/config"
	"github.com/damocles217/user_service/src/user/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(user models.UserCreate, id primitive.ObjectID, collection *mongo.Collection) (map[string]any, bool) {
	userUpdated := models.User{}
	var update primitive.D

	var inInterface map[string]interface{}

	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}

	if user.ChangePassword == "" {
		update = bson.D{
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "lastname", Value: user.Lastname}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "email", Value: user.Email}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "url_photo", Value: user.UrlPhoto}}},
		}
	} else {
		password, _ := config.HashPassword(user.ChangePassword)

		update = bson.D{
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "lastname", Value: user.Lastname}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "email", Value: user.Email}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "url_photo", Value: user.UrlPhoto}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "password", Value: password}}},
		}
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&userUpdated)

	if err != nil {
		println(err.Error())
		return map[string]any{
			"message": "cannot be saved",
		}, false
	}

	inrec, _ := json.Marshal(&userUpdated)
	json.Unmarshal(inrec, &inInterface)

	return inInterface, true

}
