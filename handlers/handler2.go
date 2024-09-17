package handlers

import (
	"io"
	"log"
	"net/http"
)

type Handler2 struct {
	logger *log.Logger
}

func NewHandler2(logger *log.Logger) *Handler2 {
	return &Handler2{logger}
}

func (handler2 *Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		handler2.logger.Println("Error in handler2:", err)
		http.Error(w, "Error in handler2", http.StatusInternalServerError)
		return
	}

	handler2.logger.Println("Handler 2 called, data received:", string(data))

	n, err := io.WriteString(w, "hello from there")
	if err != nil {
		handler2.logger.Println("Error in handler2:", err)
		http.Error(w, "Error in handler2", http.StatusInternalServerError)
		return
	}
	handler2.logger.Println("Handler 2 data written in response:", n)
}
