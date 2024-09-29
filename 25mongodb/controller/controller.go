package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = ""
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connection Established")

	collection = client.Database(dbName).Collection(colName)

	// collection refernace
	fmt.Println("Connection instance is Ready")
}
