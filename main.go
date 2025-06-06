package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/G4ble/Blog/routing"
)

func main() {
	http.HandleFunc("/static/", routing.StaticHandler)
	http.HandleFunc("/game", routing.GameHandler)
	http.HandleFunc("/*", routing.IndexHandler)

	GracefulListenAndServe(":900", nil)
}

func GracefulListenAndServe(addr string, handler http.Handler) {
	server := &http.Server{Addr: addr, Handler: handler}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener
		log.Fatal(err)
	}

	<-idleConnsClosed
}
