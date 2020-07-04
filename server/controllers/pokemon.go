package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jdotc2/blue-apricot/server/config"
)

func Gen struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var collecion *mongo.Collection

func gen8Collection() {
	collection = c.Collection('gen8')
}

func getGen8(c *gin.Context) {

}

func getGen8ById(c *gin.Context) {

}