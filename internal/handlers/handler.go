package handlers

import (
	"encoding/json"
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
	h.mux.HandleFunc("/rest/substr/find", h.loggingMiddleware(h.findHandler))
	h.mux.HandleFunc("/rest/email/check", h.loggingMiddleware(h.checkEmailAndIINHandler))

	return h.mux
}

func (h *Handler) findHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	maxSubstring := h.service.Finder.FindMaxSubstring(string(reqBody))

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(maxSubstring))
}

func (h *Handler) checkEmailAndIINHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	results := h.service.EmailChecker.CheckEmail(reqBody)

	jsonResult, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "json/application")
	w.Write(jsonResult)
}
