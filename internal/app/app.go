package app

import (
	"log"
	"net/http"
	"rest-api/internal/handlers"
)

func Start() {
	handler := handlers.NewHandler()

	server := http.Server{
		Addr:    ":8080",
		Handler: handler.InitRoutes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("error occured while starting server: %e", err)
		return
	}
}
