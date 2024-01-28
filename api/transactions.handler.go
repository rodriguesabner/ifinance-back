package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/rodriguesabner/ifinance-back/service"
	"net/http"
	"strconv"
	"time"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	month, err := strconv.Atoi(r.URL.Query().Get("month"))
	category := r.URL.Query().Get("category")

	query := struct {
		YEAR     int
		MONTH    int
		CATEGORY string
	}{year, month, category}

	mapClaimsUser := r.Context().Value("user").(*jwt.MapClaims)
	transactions, err := service.GetAllTransactions(ctx, mapClaimsUser, service.QueryFilter(query))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	RespondWithJSON(w, http.StatusOK, transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction service.TransactionToCreate
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mapClaimsUser := r.Context().Value("user").(*jwt.MapClaims)
	typeTransaction := r.URL.Query().Get("type") //outcome OR income

	if typeTransaction == "" {
		RespondWithError(w, http.StatusExpectationFailed, "Type transaction not found")
		return
	}

	if typeTransaction != "income" && typeTransaction != "outcome" {
		RespondWithError(w, http.StatusExpectationFailed, "Type transaction must be 'outcome' or 'income'")
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

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction service.TransactionToCreate
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mapClaimsUser := r.Context().Value("user").(*jwt.MapClaims)
	typeTransaction := r.URL.Query().Get("type") //outcome OR income
	idTransaction := chi.URLParam(r, "id")

	if typeTransaction != "income" && typeTransaction != "outcome" {
		RespondWithError(w, http.StatusExpectationFailed, "Type transaction must be 'outcome' or 'income'")
		return
	}

	transaction.ID = idTransaction
	transaction.USERID = (*mapClaimsUser)["id"].(string)

	result, err := service.UpdateTransaction(ctx, transaction)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, result)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction service.TransactionToCreate
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mapClaimsUser := r.Context().Value("user").(*jwt.MapClaims)
	typeTransaction := r.URL.Query().Get("type") //outcome OR income

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
