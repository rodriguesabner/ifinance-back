package main

import (
	"errors"
	"log"
	"net/http"
)

func StartServer(port string, handler http.Handler) {
	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %s", port)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error starting server: %v", err)
	}
}
