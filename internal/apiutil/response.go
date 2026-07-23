// Package apiutil holds small helpers shared by the api/* Vercel Function
// handlers: JSON responses and CORS-free error formatting.
package apiutil

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrMethodNotAllowed = errors.New("method not allowed")

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

type errorBody struct {
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, errorBody{Message: err.Error()})
}

// RequirePost rejects any method other than POST, mirroring the fact that
// every endpoint here used to be a google.script.run call (no GET semantics).
func RequirePost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		WriteError(w, http.StatusMethodNotAllowed, ErrMethodNotAllowed)
		return false
	}
	return true
}
