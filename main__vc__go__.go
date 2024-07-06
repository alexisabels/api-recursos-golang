package main

import (
	"net/http"
	"recursos/api"
)

func main() {
    http.HandleFunc("/", handler.Handler)
    http.ListenAndServe(":3000", nil)
}
