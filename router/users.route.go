package router

import (
	"github.com/go-chi/chi"
	"github.com/rodriguesabner/ifinance-back/api"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", api.GetUserHandler)
	router.Post("/", api.CreateUserHandler)

	return router
}
