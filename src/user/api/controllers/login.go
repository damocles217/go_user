package controllers

import (
	"net/http"

	"github.com/damocles217/user_service/src/user/api/config"
	"github.com/damocles217/user_service/src/user/app"
	"github.com/damocles217/user_service/src/user/core/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.UserLogin

		if err := c.BindJSON(&user); err != nil {
			println("Error in the bind of user: \n", err.Error())
		}

		userLog, access := app.GetUser(user.Email, "email", collection, nil)

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

		tokenString, err := config.JwtParse(userLog)
		if err == nil {
			c.SetCookie("t_user", tokenString, 3600*24*10, "/", "", false, true)
			c.SetCookie("c_user", userLog["code_auth"].(string), 3600*300, "/", "", false, true)

		} else {
			println("Error in token", err.Error())
		}
		c.SetCookie("t_user", tokenString, 3600*24*10, "/", "", false, true)
		c.SetCookie("c_user", userLog["code_auth"].(string), 3600*300, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}
