package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rodriguesabner/ifinance-back/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Transaction struct {
	ID          string    `bson:"_id"`
	USERID      string    `json:"user_id"`
	NAME        string    `json:"name"`
	PRICE       string    `json:"price"`
	CATEGORY    string    `json:"category"`
	DATE        time.Time `json:"date"`
	TYPE        string    `json:"type"`
	DESCRIPTION string    `json:"description"`
	PAID        bool      `json:"paid"`
}

type TransactionToCreate struct {
	ID          string    `bson:"_id"`
	USERID      string    `json:"user_id"`
	NAME        string    `json:"name"`
	PRICE       string    `json:"price"`
	CATEGORY    string    `json:"category"`
	DATE        time.Time `bson:"date"`
	TYPE        string    `json:"type"`
	DESCRIPTION string    `json:"description"`
	PAID        bool      `json:"paid"`
}

type QueryFilter struct {
	YEAR     int    `json:"year"`
	MONTH    int    `json:"month"`
	CATEGORY string `json:"category"`
}

func GetAllTransactions(ctx context.Context, mapClaimsUser *jwt.MapClaims, queryFilter QueryFilter) ([]Transaction, error) {
	collection := database.GetCollection("transactions")

	userId := (*mapClaimsUser)["id"].(string)
	andConditions := bson.A{}

	if queryFilter.MONTH != 0 {
		andConditions = append(andConditions, bson.D{{"$eq", bson.A{bson.D{{"$month", "$date"}}, queryFilter.MONTH}}})
	}
	if queryFilter.YEAR != 0 {
		andConditions = append(andConditions, bson.D{{"$eq", bson.A{bson.D{{"$year", "$date"}}, queryFilter.YEAR}}})
	}

	if queryFilter.CATEGORY != "" {
		andConditions = append(andConditions, bson.D{{"$eq", bson.A{"$category", queryFilter.CATEGORY}}})
	}

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"userid", userId},
			{"$expr", bson.D{
				{"$and", andConditions},
			}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []Transaction
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func CreateTransaction(ctx context.Context, transaction TransactionToCreate) (*mongo.InsertOneResult, error) {
	collection := database.GetCollection("transactions")
	transaction.ID = uuid.New().String()

	result, err := collection.InsertOne(ctx, transaction)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateTransaction(ctx context.Context, transaction TransactionToCreate) (*mongo.UpdateResult, error) {
	collection := database.GetCollection("transactions")

	filter := bson.M{"_id": transaction.ID}

	update := bson.M{"$set": bson.M{
		"name":        transaction.NAME,
		"price":       transaction.PRICE,
		"category":    transaction.CATEGORY,
		"date":        transaction.DATE,
		"description": transaction.DESCRIPTION,
		"paid":        transaction.PAID,
	}}
	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteTransaction(ctx context.Context, idTransaction string) (*mongo.DeleteResult, error) {
	collection := database.GetCollection("transactions")

	filter := bson.M{"_id": idTransaction}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}
