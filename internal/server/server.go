package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", Port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      setupRoutes(),
	}

	return &App{
		Server: server,
	}
}

func (app *App) Start() {
	fmt.Printf("listening on port %s\n", app.Server.Addr)
	err := app.Server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
