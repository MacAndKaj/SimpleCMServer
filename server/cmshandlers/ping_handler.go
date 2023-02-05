package cmshandlers

import (
	"log"
	"net/http"
)

type PingHandler struct {
	logger *log.Logger
}

func NewPingHandler(l *log.Logger) *Handler {
	return &Handler{
		Pattern:        "/ping",
		PatternHandler: &PingHandler{l},
		Method:         "GET",
	}
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Server ping from ", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
}
