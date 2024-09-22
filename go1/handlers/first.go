package handlers

import (
	"io"
	"log"
	"net/http"
)

type First struct {
	logger *log.Logger
}

func NewFirst(logger *log.Logger) *First {
	return &First{logger}
}

func (handler *First) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		handler.logger.Println("Error in 'First' handler:", err)
		http.Error(w, "Error in 'First' handler", http.StatusInternalServerError)
		return
	}

	handler.logger.Printf("'First' handler called, data received: \"%s\"", string(data))

	responseString := "hello"
	n, err := io.WriteString(w, responseString)
	if err != nil {
		handler.logger.Println("Error in 'First' handler:", err)
		http.Error(w, "Error in 'First' handler", http.StatusInternalServerError)
		return
	}
	handler.logger.Printf("Sent response (%d bytes): \"%s\"", n, responseString)
}
