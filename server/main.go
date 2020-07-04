package main

import (
	"log"

	"github.com/gin-gonic/gin"
	
	"github.com/jdotc2/blue-apricot/server/config"
	"github.com/jdotc2/blue-apricot/server/routes"
)

func main() {

	config.Connect()

	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}