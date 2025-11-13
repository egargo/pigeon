# pigeon

[![Go Report Card](https://goreportcard.com/badge/github.com/egargo/pigeon)](https://goreportcard.com/report/github.com/egargo/pigeon)
[![Go Reference](https://pkg.go.dev/badge/github.com/egargo/pigeon.svg)](https://pkg.go.dev/github.com/egargo/pigeon)

An simple package to manage HTTP request / response payload.


## Installation

```sh
go get github.com/egargo/pigeon
```


## Quick Start

```go
package main

import (
    "net/http"

    "github.com/egargo/pigeon"
)

type User struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Get JSON request body
    body, err := pigeon.DecodeJSON(User{}, r.Body)
    if err != nil {
        pigeon.JSON(w, http.StatusBadRequest, pigeon.S{
            Error: err.Error(),
        })
        return
    }

    /// ...

    // Return JSON response
    pigeon.JSON(w, http.StatusOK, pigeon.S{
        Message: "logged in",
    })
})
```


## License

This program is provided under the [MIT License](./LICENSE).
