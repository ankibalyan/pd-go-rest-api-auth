package main

import "pdauth/internal/server"

func main() {
	app := server.NewApp()
	app.Start()
}
