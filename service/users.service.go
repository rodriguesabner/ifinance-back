package service

import (
	"context"
	"github.com/rodriguesabner/ifinance-back/database"
	"github.com/rodriguesabner/ifinance-back/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	EMAIL    string             `bson:"email"`
	PASSWORD string             `bson:"password"`
}

type UserTokenResponse struct {
	TOKEN string `json:"token"`
}

func LoginUser(ctx context.Context, user User) (UserTokenResponse, error) {
	collection := database.GetCollection("users")
	filter := bson.M{"email": user.EMAIL, "password": user.PASSWORD}

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		log.Println(err)
		return UserTokenResponse{}, err
	}

	userId := user.ID.Hex()
	token, _ := middleware.GenerateJWT(user.EMAIL, userId)

	response := UserTokenResponse{
		TOKEN: token,
	}

	return response, err
}

func GetAllUsers(ctx context.Context) ([]User, error) {
	collection := database.GetCollection("users")
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

func CreateUser(ctx context.Context, user User) (UserTokenResponse, error) {
	collection := database.GetCollection("users")
	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		log.Fatal(err)
		return UserTokenResponse{}, err
	}

	userId := result.InsertedID.(primitive.ObjectID).Hex()
	token, _ := middleware.GenerateJWT(user.EMAIL, userId)

	response := UserTokenResponse{
		TOKEN: token,
	}

	return response, err
}
