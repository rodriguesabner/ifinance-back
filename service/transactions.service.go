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

type DatesFilter struct {
	YEAR  int `json:"year"`
	MONTH int `json:"month"`
}

func GetAllTransactions(ctx context.Context, mapClaimsUser *jwt.MapClaims, dates DatesFilter) ([]Transaction, error) {
	collection := database.GetCollection("transactions")

	userId := (*mapClaimsUser)["id"].(string)

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"userid", userId},
			{"$expr", bson.D{
				{"$and", bson.A{
					bson.D{{"$eq", bson.A{bson.D{{"$month", "$date"}}, dates.MONTH}}},
					bson.D{{"$eq", bson.A{bson.D{{"$year", "$date"}}, dates.YEAR}}},
				}},
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
