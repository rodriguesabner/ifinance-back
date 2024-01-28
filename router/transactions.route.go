package router

import (
	"github.com/go-chi/chi"
	"github.com/rodriguesabner/ifinance-back/api"
	"github.com/rodriguesabner/ifinance-back/middleware"
)

func TransactionsRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(router chi.Router) {
		router.Use(middleware.JWTMiddleware)
		router.Get("/", api.GetAllTransactions)
		router.Post("/", api.CreateTransaction)
		router.Patch("/{id}", api.UpdateTransaction)
		router.Delete("/{id}", api.DeleteTransaction)
	})

	return router
}
