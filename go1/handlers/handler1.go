package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Handler1 struct {
	logger *log.Logger
}

func NewHandler1(logger *log.Logger) *Handler1 {
	return &Handler1{
		logger: logger,
	}
}

func (h *Handler1) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Got data: %s", data)

	n, err := fmt.Fprintf(rw, "Hello, got message '%s'", string(data))
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println(n, "bytes written in response")
}
