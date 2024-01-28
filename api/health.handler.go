package api

import (
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message int    `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "healthy",
		Message: http.StatusOK,
	}

	RespondWithJSON(w, http.StatusOK, response)
}
