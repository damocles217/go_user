package user

import (
	"github.com/damocles217/user_service/src/middlewares"
	"github.com/damocles217/user_service/src/user/api/controllers"
	"github.com/damocles217/user_service/src/user/core/database"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {

	r := gin.New()

	collection := database.Connection()

	r.Use(middlewares.CORSMiddleware())

	user := r.Group("/user")
	{
		user.POST("/", controllers.CreateUser(collection))
	}

	return r
}
