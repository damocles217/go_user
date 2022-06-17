package controllers

import (
	"net/http"

	"github.com/damocles217/user_service/src/user/app"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("userId")
		filter := bson.D{
			{"name", 1},
			{"lastname", 1},
			{"gender", 1},
			{"userId", 1},
		}

		opts := options.FindOne().SetProjection(filter)

		user, _ := app.GetUser(param, "userId", collection, opts)

		c.JSON(http.StatusOK, user)
		return
	}
}
