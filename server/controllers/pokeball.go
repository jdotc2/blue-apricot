package controllers

import (
	"net/http"
	"log"

	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// BuyLocation Object
type BuyLocation struct {
	Name						string						`json:"name"`
	Location				string						`json:"location"`
	Cost						int								`json:"cost"`
	Description			string						`json:"description"`
}

// NPCLocation Object
type NPCLocation struct {
	Name						string						`json:"name"`
	Location				string						`json:"location"`
	Description			string						`json:"description"`
}

// Pokeball Object
type Pokeball struct {
	ID							int								`json:"id"`
	Name						string						`json:"name"`
	Icon						string						`json:"icon"`
	Effect					string						`json:"effect"`
	BuyLocation			BuyLocation
	NPCLocation			NPCLocation
}

// DATABASE INSTANCE
var pokeballCollection *mongo.Collection

// PokeballCollection Client Connection
func PokeballCollection(c *mongo.Database) {
	pokeballCollection = c.Collection("pokeball")
}

func (pokeball *Pokeball) setID(ID int) {
	pokeball.ID = ID
}

// GetAllPokeballs Data
func GetAllPokeballs(c *gin.Context) {
	pokeballs := []Pokeball{}
	cursor, err := pokeballCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting pokeballs, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	id := 0

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var pokeball Pokeball
		cursor.Decode(&pokeball)
		pokeball.setID(id)
		id++
		pokeballs = append(pokeballs, pokeball)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 	http.StatusOK,
		"message": 	"All Pokeballs",
		"data": 		pokeballs,
	})
	return
}