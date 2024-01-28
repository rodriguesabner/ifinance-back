package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/rodriguesabner/ifinance-back/api"
)

func configRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	return router
}

func SetupRouter() *chi.Mux {
	router := configRouter()

	apiRouter := chi.NewRouter()
	router.Mount("/v1", apiRouter)

	apiRouter.Get("/health", api.HealthHandler)
	apiRouter.Mount("/user", UserRoutes())

	return router
}
