package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/Tatsuemon/anony/usecase"
)

type HttpHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	usecase.AnonyURLUseCase
}

func NewHttpHandler(u usecase.AnonyURLUseCase) HttpHandler {
	return &httpHandler{u}
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("SERVER_HOST")
	anURL := host + r.URL.String()

	original, err := h.AnonyURLUseCase.GetOriginalByAnonyURL(context.Background(), anURL)
	if err != nil {
		// TODO(Tatsuemon): errorの時の処理
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", original)
	// 301 Moved Permanently
	w.WriteHeader(http.StatusMovedPermanently)
	return
}
