package handlers

import (
	"log"
	"net/http"
)

type RootHandler struct {
	logger *log.Logger
}

func NewRoot(l *log.Logger) *RootHandler {
	return &RootHandler{l}
}

func (h *RootHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	log.Println("Attempting root site")

}
