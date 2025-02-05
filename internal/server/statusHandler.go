package server

import (
	"fmt"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request accepted at status handler.")
	fmt.Fprint(w, "OK")
}
