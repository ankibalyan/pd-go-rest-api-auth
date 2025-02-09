package httputils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrEmptyBody = errors.New("EMPTY_BODY")

func ParseRequestBody(r *http.Request, result any) error {
	reqBody := r.Body
	defer reqBody.Close()

	if reqBody == nil {
		return ErrEmptyBody
	}

	decoder := json.NewDecoder(reqBody)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(result)

	if err != nil {
		return fmt.Errorf("failed in parsing request body: %w", err)
	}

	return nil
}

func SendJSONResponse(w http.ResponseWriter, data any) {
	// response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)

	if err != nil {
		http.Error(w, "something went wrong!", http.StatusInternalServerError)

		return
	}
}
