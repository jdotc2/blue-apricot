package main

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jdotc2/blue-apricot/config"
	"github.com/jdotc2/blue-apricot/routes"
)

func main() {
	fmt.Println("Sup World")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}