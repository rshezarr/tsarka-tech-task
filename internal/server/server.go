package server

import (
	"context"
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
	Srv               *http.Server
	ServerErrorNotify chan error
}

func NewServer(router *http.ServeMux) *Server {
	return &Server{
		Srv: &http.Server{
			Addr:           Addr,
			Handler:        router,
			MaxHeaderBytes: MaxHeaderBytes,
			ReadTimeout:    ReadTimeout,
			WriteTimeout:   WriteTimeout,
		},
		ServerErrorNotify: make(chan error, 1),
	}
}

func (s *Server) Start() {
	s.ServerErrorNotify <- s.Srv.ListenAndServe()
}

func (s *Server) ServerErrNotify() <-chan error {
	return s.ServerErrorNotify
}

func (s *Server) Shutdown() error {
	return s.Srv.Shutdown(context.Background())
}
