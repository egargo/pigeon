package pigeon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// H
type H struct {
	Message string         `json:"message,omitempty"`
	Error   string         `json:"error,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, r *http.Request, status int, data H) {
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

func DecodeJSON[T any](r *http.Request, class T, data io.Reader) (T, error) {
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
