package server

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Println("Request accepted at index handler.")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome to planet dev."))
}
