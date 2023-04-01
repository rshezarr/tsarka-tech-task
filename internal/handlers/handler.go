package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rest-api/internal/service"
)

type Handler struct {
	mux     *http.ServeMux
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		mux:     http.NewServeMux(),
		service: service,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	h.mux.HandleFunc("/rest/substr/find", h.findHandler)
	h.mux.HandleFunc("/rest/email/check", h.checkEmailAndIINHandler)

	return h.mux
}

func (h *Handler) findHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}
	s, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	maxSubstring := h.service.Finder.FindMaxSubstring(string(s))
	fmt.Fprint(w, maxSubstring)
}

func (h *Handler) checkEmailAndIINHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}
	s, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	results := h.service.EmailChecker.CheckEmail(s)

	jsonResult, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "json/application")
	w.Write(jsonResult)
}
