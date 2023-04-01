package app

import (
	"log"
	"net/http"
	"rest-api/internal/handlers"
	"rest-api/internal/service"
)

func Start() {
	service := service.NewService()
	handler := handlers.NewHandler(service)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler.InitRoutes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("error occured while starting server: %e", err)
		return
	}
}
