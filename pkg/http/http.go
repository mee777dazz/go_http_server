package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"http_server/cmd/app/config"
)

func CreateAndRunServer(r chi.Router, cfg config.HTTPConfig) error {
	httpServer := &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
