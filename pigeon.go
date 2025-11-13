package pigeon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// S is a predefined structure.
// Message, for information about the response.
// Error, for error information about the response.
// Data, for data that will be return with the response.
type S struct {
	Message string         `json:"message,omitempty"`
	Error   string         `json:"error,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
}

// H is a shortcut for map[string]any, used for free-form structure.
type H map[string]any

// DecodeJSON decodes a JSON payload from the given io.Reader into a value of type T.
func DecodeJSON[T any](class T, data io.Reader) (T, error) {
	var m any
	mErr := json.NewDecoder(data).Decode(&m)
	if mErr != nil {
		return class, mErr
	}

	b, bErr := json.Marshal(m)
	if bErr != nil {
		return class, bErr
	}

	var t T
	tErr := json.Unmarshal(b, &t)
	if tErr != nil {
		return class, tErr
	}

	return t, nil
}

// JSON writes the given data as a JSON response with the specified HTTP status code.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8;")

	var response bytes.Buffer

	err := json.NewEncoder(&response).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}

	w.WriteHeader(status)
	w.Write(response.Bytes())
}
