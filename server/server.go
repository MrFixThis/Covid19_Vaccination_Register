package server

import (
	"net/http"
	"time"
)

// InitializeServer initializes the server instance
func InitializeServer(r http.Handler) *http.Server {
	return &http.Server{
		Addr:              "127.0.0.1:9000",
		Handler:           r,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	// Read about all the field of the http.Server's fields, and the suggested
	// values for each of them
}
