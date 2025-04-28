package utils

import (
	"encoding/json"
	"net/http"
)

// Pagination metadata
type Pagination struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

// Envelope response
type Envelope struct {
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

// Use this for errors
func WriteJSONError(w http.ResponseWriter, status int, message string) error {
	env := Envelope{
		Error: message,
	}
	return WriteJSON(w, status, env)
}

// Use this for normal data responses
func JsonResponse(w http.ResponseWriter, status int, data any) error {
	env := Envelope{
		Data: data,
	}
	return WriteJSON(w, status, env)
}

// Use this for paginated data responses
func JsonResponseWithPagination(w http.ResponseWriter, status int, data any, pagination Pagination) error {
	env := Envelope{
		Data:       data,
		Pagination: &pagination,
	}
	return WriteJSON(w, status, env)
}
