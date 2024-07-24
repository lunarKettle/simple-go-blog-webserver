package middleware

import (
	"log"
	"net/http"
	"time"
)

func TimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		milliseconds := float64(duration) / float64(time.Millisecond)
		log.Printf("Request %s %s took %.2f ms", r.Method, r.RequestURI, milliseconds)
	})
}
