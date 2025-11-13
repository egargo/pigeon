package main

import (
	"net/http"

	"github.com/egargo/pigeon"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	http.ListenAndServe(":8080", mux)
}
