package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from the panic, %s", err)
				http.Error(w, "Internal server error.", http.StatusInternalServerError)
			}
		}()

		fmt.Println("recover middleware.")
		next.ServeHTTP(w, r)
		fmt.Println("recover middleware end.")
	})
}
