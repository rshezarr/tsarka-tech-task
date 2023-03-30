package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
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
	h.mux.HandleFunc("/rest/substr/find", h.findHandler)
	h.mux.HandleFunc("/rest/email/check", h.checkEmailAndIINHandler)

	return h.mux
}

func findMaxSubstring(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		sub := ""
		for j := i; j < len(s); j++ {
			if index := strings.IndexByte(sub, s[j]); index == -1 {
				sub += string(s[j])
			} else {
				break
			}
		}
		if utf8.RuneCountInString(sub) > utf8.RuneCountInString(result) {
			result = sub
		}
	}
	return result
}

func (h *Handler) findHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	maxSubstring := findMaxSubstring(string(s))
	fmt.Fprint(w, maxSubstring)
}

func (h *Handler) checkEmailAndIINHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	emailRegex := regexp.MustCompile(`(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`)
	iinRegex := regexp.MustCompile(`^[0-9]{12}$`)

	emails := emailRegex.FindAllString(string(s), -1)
	iins := iinRegex.FindAllString(string(s), -1)

	results := make(map[string]interface{})
	results["emails"] = emails
	results["iins"] = iins

	jsonResult, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "json/application")
	w.Write(jsonResult)
}
