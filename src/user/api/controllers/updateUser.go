package controllers

import (
	"net/http"

	"github.com/damocles217/user_service/src/user/api/config"
	"github.com/damocles217/user_service/src/user/app"
	"github.com/damocles217/user_service/src/user/core/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.UserCreate

		if err := c.BindJSON(&user); err != nil {
			println("Error in the bind of user: \n", err.Error())
		}

		tokenString, _ := c.Cookie("t_user")
		id := config.GetUserId(tokenString)

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

		userUpdated, accessTo := app.UpdateUser(user, id, collection)

		if accessTo {
			c.JSON(http.StatusOK, userUpdated)
			return
		}

		c.JSON(http.StatusBadRequest, userUpdated)
		return
	}
}
