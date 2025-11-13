package main

import (
	"net/http"

	"github.com/egargo/pigeon"
)

type Hello struct {
	Name string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	pigeon.JSON(w, http.StatusOK, pigeon.H{
		"hello": "world",
	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	body, err := pigeon.DecodeJSON(Hello{}, r.Body)
	if err != nil {
		pigeon.JSON(w, http.StatusBadRequest, pigeon.S{
			Error: err.Error(),
		})
		return
	}

	pigeon.JSON(w, http.StatusOK, pigeon.H{
		"hello": body.Name,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", helloWorldHandler)
	mux.HandleFunc("POST /", greetHandler)

	http.ListenAndServe(":2122", mux)
}
