package api

import (
	"context"
	"encoding/json"
	"github.com/rodriguesabner/ifinance-back/service"
	"net/http"
	"time"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := service.GetAllUsers(ctx)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	RespondWithJSON(w, http.StatusOK, users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user service.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := service.CreateUser(ctx, user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, result)
}
