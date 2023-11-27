package controllers

import (
	"net/http"
	"log"
	"time"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "github.com/jdotc2/blue-apricot/server/config"
)

// Gen8 Object
type Gen8 struct {
	ID        string    	`json:"id"`
	Name	  string		`json:"name"`
	Images	  struct {
		Main	string		`json:"main"`
		Map		string		`json:"map"`
	}
	Exclusive struct {
		Pokemon	[]int		`json:"pokemon"`
	}
	Completed string    	`json:"completed"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}


// DATABASE INSTANCE
var collection *mongo.Collection

// Gen8Collection Client Connection
func Gen8Collection(c *mongo.Database) {
	collection = c.Collection("gen8")
}

// GetAllGen8 Get all
func GetAllGen8(c *gin.Context) {
	gen8s := []Gen8{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
			var gen8 Gen8
			cursor.Decode(&gen8)
			gen8s = append(gen8s, gen8)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Gen8s",
		"data":    gen8s,
	})
	return
}
