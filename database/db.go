package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// GetCollection returns a collection from the database
func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database("test").Collection(name)
}
// StartMongoDB starts the mongodb database
func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	return nil
}

// CloseMongoDB closes the mongodb database
func CloseMongoDB() error {
	err := mongoClient.Disconnect(context.TODO());
	if err != nil {
		panic(err)
	}
	return nil
}
