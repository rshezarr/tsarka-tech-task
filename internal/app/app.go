package app

import (
	"log"
	"os"
	"os/signal"
	"rest-api/internal/handlers"
	"rest-api/internal/server"
	"rest-api/internal/service"
	"syscall"
)

func Start() {
	service := service.NewService()
	handler := handlers.NewHandler(service)

	srv := server.NewServer(handler.InitRoutes())

	log.Printf("app: starting...")

	quit := make(chan os.Signal, 1)

	go func() {
		log.Printf("app: Starting server at port %v -> http://localhost%v", srv.Srv.Addr, srv.Srv.Addr)
		srv.Start()
	}()

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	select {
	case signal := <-quit:
		log.Printf("app: signal accepted: %v", signal)
	case err := <-srv.ServerErrNotify():
		log.Printf("app: server closing: %v", err)
	}

	log.Printf("app: shutting down...")
}
