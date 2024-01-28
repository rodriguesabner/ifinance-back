package service

import (
	"context"
	"github.com/rodriguesabner/ifinance-back/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct {
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

var collection *mongo.Collection = database.GetCollection("users")

func GetAllUsers(ctx context.Context) ([]User, error) {
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []User
	for cursor.Next(ctx) {
		var document User
		if err := cursor.Decode(&document); err != nil {
			log.Fatal(err)
		}
		results = append(results, document)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func CreateUser(ctx context.Context, user User) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	return result, nil
}
