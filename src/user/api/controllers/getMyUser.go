package controllers

import (
	"net/http"
	"os"

	"github.com/damocles217/user_service/src/user/app"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMyUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.MapClaims{
			"_id":       "",
			"code_auth": "",
		}

		token, err := c.Cookie("t_user")
		cUser, err := c.Cookie("c_user")

		var id string

		if err != nil {
			println("Error, cookie not found\n ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No credentials found",
			})
			return
		}

		secret := os.Getenv("SECRET_JWT")

		_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		for key, val := range claims {
			if key == "_id" {
				id = val.(string)
			}
		}

		idObj, err := primitive.ObjectIDFromHex(id)

		data, _ := app.GetUser(idObj, "_id", collection, nil)

		cAuth := data["code_auth"]

		if cUser == cAuth {
			c.JSON(http.StatusOK, data)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This is not your profile",
		})
		return
	}
}
