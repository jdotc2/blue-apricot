package test

import (
	"context"
)

type Test struct {
	ID string "id"
}

func GetTest(c *gin.Context) {
	test := []Test{}
	log.Printf(test)
}