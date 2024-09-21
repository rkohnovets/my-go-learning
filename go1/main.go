package main

import (
	"context"
	"go1/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	consoleLogger := log.New(os.Stdout, "console log", log.Ldate|log.Ltime)

	firstHandler := handlers.NewFirst(consoleLogger)
	productsHandler := handlers.NewProducts(consoleLogger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/first", firstHandler)
	serveMux.Handle("/products", productsHandler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			consoleLogger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	consoleLogger.Println("Received terminate, gracefully shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(tc)
	if err != nil {
		consoleLogger.Fatal("shutting down error", err)
	}
}
