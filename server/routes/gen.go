package routes

import (
	"github.com/gin-gonic/gin"

	// Controllers
	"github.com/jdotc2/blue-apricot/server/controllers"
)

func genRoutes(router *gin.Engine) {
	router.GET("/gen8", controllers.GetAllGen8)
}