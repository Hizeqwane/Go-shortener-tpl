package handler

import (
	"github.com/Hizeqwane/Go-shortener-tpl/internal/service"
	"io"
	"net/http"
	"strings"
)

type RequestHandler struct {
	shortenerService service.IShortenerService
}

func NewHandler() *RequestHandler {
	return &RequestHandler{
		shortenerService: service.NewShortenerService(),
	}
}

func (p *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.Header.Get("Content-Type") != "text/plain" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		longUri := string(bodyBytes)

		shortUri := p.shortenerService.GetShortUri(longUri)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortUri))
	}

	if r.Method == http.MethodGet {
		shortUri := strings.Trim(r.URL.Path, "/")

		if i, ok := p.shortenerService.TryGetLongUri(shortUri); ok {
			w.Header().Set("Location", i)
			w.WriteHeader(http.StatusTemporaryRedirect)

			return
		}

		w.WriteHeader(http.StatusBadRequest)
	}
}

var _ http.Handler = (*RequestHandler)(nil)
