package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/rodriguesabner/ifinance-back/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id"`
	USERID      string             `json:"user_id"`
	NAME        string             `json:"name"`
	PRICE       string             `json:"price"`
	CATEGORY    string             `json:"category"`
	DATE        time.Time          `json:"date"`
	TYPE        string             `json:"type"`
	DESCRIPTION string             `json:"description"`
	PAID        bool               `json:"paid"`
}

func GetAllFinances(ctx context.Context, mapClaimsUser *jwt.MapClaims) ([]Transaction, error) {
	collection := database.GetCollection("transactions")

	userId := (*mapClaimsUser)["id"].(string)
	filter := bson.M{"userid": userId}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []Transaction
	for cursor.Next(ctx) {
		var document Transaction
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

func CreateTransaction(ctx context.Context, transaction Transaction) (*mongo.InsertOneResult, error) {
	collection := database.GetCollection("transactions")
	result, err := collection.InsertOne(ctx, transaction)

	if err != nil {
		return nil, err
	}

	return result, nil
}
