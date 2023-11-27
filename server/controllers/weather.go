package controllers

import (
	"net/http"
	"log"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "github.com/jdotc2/blue-apricot/server/config"
)

// Weather Object
type Weather struct {
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	Sprite			string		`json:"sprite"`
	Date			string		`json:"date"`
}

// DATABASE INSTANCE
var weatherCollection *mongo.Collection

// WeatherCollection Client Connection
func WeatherCollection(c *mongo.Database) {
	weatherCollection = c.Collection("weather")
}

// GetAllWeather Data
func GetAllWeather(c *gin.Context) {
	weathers := []Weather{}
	cursor, err := weatherCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all weather, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Someting went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var weather Weather
		cursor.Decode(&weather)
		weathers = append(weathers, weather)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 	http.StatusOK,
		"message": 	"All Weathers",
		"data": 		weathers,
	})
	return
}