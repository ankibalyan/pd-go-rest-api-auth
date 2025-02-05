package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)

		log.Printf("%s %s Duration: %s", r.Method, r.URL, time.Since(start))
	})
}
