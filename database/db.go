package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

const DbName = "ifinance"

func ConnectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURL := os.Getenv("PORT")

	opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)

	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := Client.Database(DbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func GetCollection(CollectionName string) *mongo.Collection {
	return Client.Database(DbName).Collection(CollectionName)
}
