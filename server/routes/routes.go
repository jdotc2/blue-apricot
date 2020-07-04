package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routes Function
func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.NoRoute(notFound)
	router.GET("/gen8/:name/", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
