package controller

import (
	"25mongodb/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://tanwar0210:YteungvSYaO25A1C@cluster0.6pb4v.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
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

// Mongo helpers -> File

// insert 1 record
func insertOneMovie(movie model.Netflix){
	success, error := collection.InsertOne(context.Background(), movie)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Inserted One Movie", success.InsertedID)
}

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated:", result.ModifiedCount)
	
}