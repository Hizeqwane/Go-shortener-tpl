package main

import (
	"github.com/Hizeqwane/Go-shortener-tpl/internal/handler"
	"github.com/Hizeqwane/Go-shortener-tpl/internal/service"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// создаём один экземпляр сервиса
	shortenerService := service.NewShortenerService()

	// хендлер использует этот же сервис
	shortenerHandler := handler.NewHandler(shortenerService)

	mux.Handle("/", shortenerHandler)

	http.ListenAndServe(":8080", mux)
}
