package server

import (
	"net/http"
	"pdauth/internal/domain/auth"
	"pdauth/internal/middlewares"
)

func setupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", indexHandler)
	mux.Handle("POST /status", http.HandlerFunc(statusHandler))

	mux.HandleFunc("POST /auth/signup", auth.SignupHanlder)

	handler := middlewares.RecoverMiddleware(
		middlewares.LoggerMiddleware(mux),
	)

	return handler
}
