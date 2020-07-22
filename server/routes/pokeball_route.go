package routes

import (
	"github.com/gin-gonic/gin"

	// Controllers
	"github.com/jdotc2/blue-apricot/server/controllers"
)

// PokeballRoutes Gets all Pokeball metadata
func PokeballRoutes(router *gin.Engine) {
	router.GET("/weather", controllers.GetAllPokeballs)
}