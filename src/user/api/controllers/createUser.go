package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/damocles217/user_service/src/user/app"
	"github.com/damocles217/user_service/src/user/core/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.UserCreate

		if err := c.BindJSON(&user); err != nil {
			println("Error in the bind of user: \n", err.Error())
		}

		if user.ConfirmPassword != user.Password {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Password has to match with Confirm password",
			})
			return
		}

		// Setting the values for being correctly saved
		userForCreating := models.User{
			Name:      user.Name,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Password:  user.Password,
			Gender:    user.Gender,
			CreatedAt: time.Now(),
			Admin:     0,
		}
		userForCreating.UpdatedAt = userForCreating.CreatedAt

		mapping, access := app.CreateUser(userForCreating, collection)

		if access {
			// JWT usage
			// Declare secret
			secret := os.Getenv("SECRET_JWT")

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"_id": mapping["_id"],
			})

			// String token
			tokenString, err := token.SignedString([]byte(secret))
			if err == nil {
				c.SetCookie("t_user", tokenString, 3600*24*10, "/", "", false, true)
				c.SetCookie("c_user", mapping["code_auth"], 3600*300, "/", "", false, true)

			} else {
				println("Error in token", err.Error())
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})

		return
	}
}
