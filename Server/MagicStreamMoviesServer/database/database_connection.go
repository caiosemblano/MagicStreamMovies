package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() *mongo.Client{
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: unable to load .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")

	if MongoDb == "" {
		log.Fatal("Error: MONGODB_URI not set!")
	}

	fmt.Println("MongoDB URI:", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)
	
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil
	}

	return client
}

var Client *mongo.Client = Connect()

func OpenCollection(collectionName string ) *mongo.Collection{
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to load .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("Error: DATABASE_NAME not set!")
	}
	
	fmt.Println("Database Name:", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}
	return collection
}	