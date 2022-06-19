package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	var cookies [2]string
	value := []string{"c_user", "t_user"}
	var err error = nil

	for i := 0; i < 2; i++ {
		cookies[i], err = c.Cookie(value[i])

		if err != nil {
			cookies[i] = "NotSet"
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "You need to be loged",
			})
			return
		}
	}

	c.SetCookie("c_user", "", -1, "/", "", false, true)
	c.SetCookie("t_user", "", -1, "/", "", false, true)
	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
	})
	return
}
