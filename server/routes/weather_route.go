package routes

import (
	"github.com/gin-gonic/gin"

	// Controllers
	"github.com/jdotc2/blue-apricot/server/controllers"
)

func weatherRoutes(router *gin.Engine) {
	router.GET("/weather", controllers.GetAllWeather)
}