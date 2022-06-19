package config

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserId(token string) primitive.ObjectID {
	var id string = ""
	claims := jwt.MapClaims{
		"_id": "",
	}

	secret := os.Getenv("SECRET_JWT")

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	for key, val := range claims {
		if key == "_id" {
			id = val.(string)
		}
	}

	idObj, _ := primitive.ObjectIDFromHex(id)
	return idObj
}
