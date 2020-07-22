package routes

import (
	"github.com/gin-gonic/gin"

	// Controllers
	"github.com/jdotc2/blue-apricot/server/controllers"
)

// TypeRoutes Grabs all pokemon type metadata
func TypeRoutes(router *gin.Engine) {
	router.GET("/weather", controllers.GetAllTypes)
}