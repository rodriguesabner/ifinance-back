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
		router.Get("/", api.GetAllFinances)
		router.Post("/", api.CreateTransaction)
	})

	return router
}
