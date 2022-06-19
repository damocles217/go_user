package controllers

import (
	"net/http"
	"strconv"

	"github.com/damocles217/user_service/src/user/app"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		page := c.Param("page")

		if page == "" {
			page = "0"
		}

		pageInt, _ := strconv.ParseInt(page, 10, 64)

		users, access := app.GetUsers(collection, pageInt)

		if !access {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "An Error ocurrent",
			})
			return
		}

		c.JSON(http.StatusOK, users)
		return
	}
}
