package api

import (
	"context"
	"encoding/json"
	"github.com/rodriguesabner/ifinance-back/service"
	"net/http"
	"time"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user service.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.LoginUser(ctx, user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid Credentials")
		return
	}

	RespondWithJSON(w, http.StatusOK, response)
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
