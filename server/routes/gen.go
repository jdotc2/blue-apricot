package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// Controllers
	"github.com/jdotc2/blue-apricot/server/controllers"
)

func genRoutes(router *gin.Engine) {
	router.GET("/gen8/:name", findGen8)
	
}

func findGen8(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, )
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}