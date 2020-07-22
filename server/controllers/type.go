package controllers

import (
	"net/http"
	"log"

	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// PType Object
type PType struct {
	ID							int								`json:"id"`
	Name						string						`json:"name"`
	NotEffective		[]string					`json:"notEffective"`
	NoEffect				[]string					`json:"noEffect"`
	VeryEffective		[]string					`json:"veryEffective"`
	Notes						string						`json:"notes"`
}

// DATABASE INSTANCE
var typeCollection *mongo.Collection

// TypeCollection Client Connection
func TypeCollection(c *mongo.Database) {
	typeCollection = c.Collection("type")
}

func (ptype *PType) setID(ID int) {
	ptype.ID = ID
}

// GetAllTypes Data
func GetAllTypes(c *gin.Context) {
	ptypes := []PType{}
	cursor, err := typeCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting types, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	id := 0

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var ptype PType
		cursor.Decode(&ptype)
		ptype.setID(id)
		id++
		ptypes = append(ptypes, ptype)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 	http.StatusOK,
		"message": 	"All Types",
		"data": 		ptypes,
	})
	return
}