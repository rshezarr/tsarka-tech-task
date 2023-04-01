package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) loggingMiddleware(router http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s [%s]\t%s%s - 200 - OK\n", time.Now().Format("2006/01/02 15:04:05"), r.Proto, r.Method, r.Host, r.RequestURI)
		router.ServeHTTP(w, r)
	}
}
