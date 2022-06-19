package app

import (
	"context"
	"strconv"

	"github.com/damocles217/user_service/src/user/core/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO create element bcrypt function for passwords
func CreateUser(user models.User, collection *mongo.Collection) (map[string]any, bool) {

	id := uuid.New()
	userId := UserId(user.Name, user.Lastname, collection)

	user.CodeAuth = id.String()
	user.UserID = userId

	newUserId, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		println(err.Error())
		return map[string]any{
			"message": "Not null",
		}, false
	}
	// Transform in valid _id primitive
	oid, ok := newUserId.InsertedID.(primitive.ObjectID)

	if ok {
		return map[string]any{
				// ID to string
				"_id":       oid.Hex(),
				"code_auth": user.CodeAuth,
			},
			true
	}

	return map[string]any{
		"message": "User cannon be saved",
	}, false
}

func UserId(name string, lastname string, coll *mongo.Collection) string {
	userId := name + "." + lastname

	for i := 0; i != -2; i++ {
		_, truth := GetUser(userId, "userId", coll, nil)
		if !truth {
			return userId
		}
		excludes := (i / 10) + 1
		userId = string([]byte(userId)[:len(userId)-excludes])
		userId += strconv.Itoa(i)
	}
	return ""
}
