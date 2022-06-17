package validators

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserValidator() primitive.M {

	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"email", "name", "lastname", "password", "gender", "userId"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "User name is necesary",
				"pattern":     "[a-zA-Z]{1}",
			},
			"lastname": bson.M{
				"bsonType":    "string",
				"description": "Laast name is necesary",
				"pattern":     "[a-zA-Z]{1}",
			},
			"email": bson.M{
				"bsonType":    "string",
				"description": "User email has to be valid",
				"pattern":     ".*@[a-zA-Z]*.com",
			},
			"password": bson.M{
				"bsonType": "string",
			},
			"userId": bson.M{
				"bsonType": "string",
			},
			"code_auth": bson.M{
				"bsonType": "string",
			},
			"admin": bson.M{
				"bsonType": "int",
				"minimum":  0,
				"maximum":  4,
			},
			"gender": bson.M{
				"bsonType":    "int",
				"enum":        []int{0, 1},
				"description": "Gender has to be between 0 and 1",
			},
			"url_photo": bson.M{
				"bsonType": "string",
				"pattern":  ".*.[png|jpe?g|gif]",
			},
			"createdAt": bson.M{
				"bsonType": "date",
			},
			"updatedAt": bson.M{
				"bsonType": "date",
			},
		},
	}

	var validator = bson.M{
		"$jsonSchema": jsonSchema,
	}
	return validator
}
