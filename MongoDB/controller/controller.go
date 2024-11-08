package controller

import (
	model "Mongo/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "MERN-Test"
const colName = "go-test"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

func InsertMovie(movie model.Netflix) {
	insert, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("Error inserting document:", err)
	}

	fmt.Println("Inserted one document with ID:", insert.InsertedID)
}
