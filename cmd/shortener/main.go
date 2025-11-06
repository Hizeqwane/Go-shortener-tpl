package main

import (
	"github.com/Hizeqwane/Go-shortener-tpl/internal/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	shortenerHandler := handler.NewHandler()

	mux.Handle("/", shortenerHandler)

	http.ListenAndServe(":8080", mux)
}
