package api

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/rodriguesabner/ifinance-back/service"
	"net/http"
	"time"
)

func GetAllFinances(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mapClaimsUser := r.Context().Value("user").(*jwt.MapClaims)
	expenses, err := service.GetAllFinances(ctx, mapClaimsUser)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	RespondWithJSON(w, http.StatusOK, expenses)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction service.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mapClaimsUser := r.Context().Value("user").(*jwt.MapClaims)
	typeTransaction := r.URL.Query().Get("type") //expense OR income

	if typeTransaction == "" {
		RespondWithError(w, http.StatusExpectationFailed, "Type transaction not found")
		return
	}

	transaction.TYPE = typeTransaction
	transaction.USERID = (*mapClaimsUser)["id"].(string)

	result, err := service.CreateTransaction(ctx, transaction)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, result)
}
