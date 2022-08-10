package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitializeServer(r *mux.Router) *http.Server {
	s := &http.Server{
		Addr:              "127.0.0.1:9000",
		Handler:           r,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Minute,
	}

	return s
}
