package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// InitializeServer initializes the server instance
func InitializeServer(r *mux.Router) *http.Server {
	return &http.Server{
		Addr:              "127.0.0.1:9000",
		Handler:           r,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Minute,
	}
	// Read about all the field of the http.Server's fields, and the suggested
	// values for each of them
}
