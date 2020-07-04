package config

import (
	"fmt"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/jdotc2/blue-apricot/server/controllers"
)

// Connect to database
func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
	}()

	db := client.Database("mew")
	controllers.gen8Collection(db)

	return
}

// Test Connection to database 
func Test() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/test?w=majority"))
		if err != nil {
			panic(err)
		}
	
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	
		// Ping the primary
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic(err)
		}
	
		fmt.Println("Successfully connected and pinged.")
}