package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

// Connection URI
const uri = "mongodb://localhost:27017/"

func InitMongoDb() {
	// Create a new client and connect to the server
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = MongoClient.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Ping the primary
	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Successfully connected and pinged.")
}
