package router

import (
	"github.com/go-chi/chi"
	"github.com/rodriguesabner/ifinance-back/api"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/login", api.LoginUserHandler)
	router.Post("/register", api.CreateUserHandler)

	return router
}
