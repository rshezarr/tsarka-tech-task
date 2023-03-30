package handlers

import (
	"net/http"
)

type Handler struct {
	mux *http.ServeMux
}

func NewHandler() *Handler {
	return &Handler{
		mux: http.NewServeMux(),
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	h.mux.HandleFunc("/rest/substr/find", nil)
	h.mux.HandleFunc("/rest/email/check", nil)

	return h.mux
}
