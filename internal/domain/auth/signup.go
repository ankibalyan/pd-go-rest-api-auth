package auth

import (
	"encoding/json"
	"net/http"
)

func SignupHanlder(w http.ResponseWriter, r *http.Request) {
	// read the request body
	// validate the request body
	// validations for signup
	// create the user
	//

	result := map[string]any{
		"status": "OK",
		"data":   nil,
		"msg":    "user created",
	}

	// response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
