package server

import (
	"net/http"
	"time"
)

const (
	Addr           = ":9090"
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
)

type Server struct {
	server *http.Server
}

func NewServer(router *http.ServeMux) *Server {
	return &Server{
		server: &http.Server{
			Addr:           Addr,
			Handler:        router,
			MaxHeaderBytes: MaxHeaderBytes,
			ReadTimeout:    ReadTimeout,
			WriteTimeout:   WriteTimeout,
		},
	}
}
