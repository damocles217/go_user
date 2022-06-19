package guards

import (
	"net/http"

	"github.com/damocles217/user_service/src/user/api/config"
	"github.com/damocles217/user_service/src/user/app"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthGuard(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("t_user")
		cUser, err := c.Cookie("c_user")

		if err != nil {
			println("Error, cookie not found\n ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No credentials found",
			})
			c.Abort()
			return
		}

		idObj := config.GetUserId(token)

		data, _ := app.GetUser(idObj, "_id", collection, nil)

		cAuth := data["code_auth"]

		if cUser == cAuth {
			c.Next()
			return
		}
		c.Abort()
	}
}
