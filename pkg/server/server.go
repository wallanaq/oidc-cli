package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var httpServer *http.Server

// Start initiates the server at the specified address and serves requests
// using the provided HTTP handler. It is a blocking call until the server
// is shut down.
//
// Example:
//
//	err := Start(":8080", handler)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//
//	addr - the address to bind the server to, e.g., ":8080"
//	handler - the HTTP handler to process incoming requests
//
// Returns:
//
//	error - if the server fails to start
func Start(addr string, handler http.Handler) error {
	httpServer = &http.Server{Addr: addr, Handler: handler}

	log.Printf("Starting server at %s\n", addr)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("ListenAndServe(): %v", err)
	}

	return nil

}

// Shutdown gracefully shuts down the server, ensuring all connections are
// properly closed before termination. It returns an error if there is an issue
// with the shutdown process.
//
// Example:
//
//	err := Shutdown()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Returns:
//
//	error - if there is an issue during server shutdown, otherwise nil
func Shutdown() error {
	if httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			return fmt.Errorf("error when shutting down the server: %v", err)
		}

		log.Println("Server gracefully stopped")
	}

	return nil
}
