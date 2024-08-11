package middleware

import (
	"log"
	"net/http"
	"time"
)

type MiddlewareService struct{}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func CheckHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for a specific header
		if r.Header.Get("X-Custom-Header") != "expected-value" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Call the next handler if headers are valid
		next.ServeHTTP(w, r)
	})
}
