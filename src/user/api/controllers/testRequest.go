package controllers

import "github.com/gin-gonic/gin"

func TestRequest(c *gin.Context) {
	c.String(200, "data")
	return
}
