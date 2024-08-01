package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(rw http.ResponseWriter, status int, v any) error {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)

	return json.NewEncoder(rw).Encode(v)
}

func WriteError(rw http.ResponseWriter, status int, err error) {
	WriteJSON(rw, status, map[string]string{"error": err.Error()})
}
