package controllers

import (
	"net/http"

	"github.com/damocles217/user_service/src/user/api/config"
	"github.com/damocles217/user_service/src/user/app"
	"github.com/damocles217/user_service/src/user/core/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.UserLogin
		c.BindJSON(&user)

		token, _ := c.Cookie("t_user")

		id := config.GetUserId(token)

		userLog, access := app.GetUser(id, "_id", collection, nil)

		if !access {
			c.JSON(http.StatusBadRequest, userLog)
			return
		}

		hashedPassword := userLog["password"].(string)

		checked := config.CheckPasswordHash(user.Password, hashedPassword)

		if !checked {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User Not Found",
			})
			return
		}

		userDel, access := app.DeleteUser(id, collection)

		if access {
			c.JSON(http.StatusOK, userDel)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not found user",
		})

	}
}
