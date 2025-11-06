package handler

import (
	"github.com/Hizeqwane/Go-shortener-tpl/internal/service"
	"io"
	"net/http"
	"strings"
)

// Глобальный сервис
var defaultService = service.NewShortenerService()

type RequestHandler struct {
	shortenerService service.IShortenerService
}

func NewHandler(shortenerService service.IShortenerService) *RequestHandler {
	return &RequestHandler{
		shortenerService: defaultService,
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

		// Если нет протокола, добавляем http://
		if !strings.HasPrefix(longUri, "http://") && !strings.HasPrefix(longUri, "https://") {
			longUri = "http://" + longUri
		}

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
