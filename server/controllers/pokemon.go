package controllers

import (
	"time"
	"github.com/gin-gonic/gin"
)

// Pokemon object
type Pokemon struct {
	ID					string			`json:"id"`
	Name				string			`json:"name"`
	CreatedAt		time.Time		`json:"created_at"`
}

// Database instance
var collection *mongo.Collection

func PokemonCollection(c *mongo.Database) {
	collection = c.Collection("mew")
}

func getAllPokemon(c *gin.Context) {

	pokemon := []Pokemon{}

	cursor, err := collection.Find(context.POKEMON(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all pokemon, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.POKEMON()) {
		var pokemon Pokemon
		cursor.Decode(&pokemon)
		allPokemon = append(allPokemon, pokemon)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "All Pokemon",
		"data": allPokemon,
	})
	return
}