package cmshandlers

import (
	"log"
	"net/http"
)

type RootHandler struct {
	logger *log.Logger
}

func NewRootHandler(l *log.Logger) *Handler {
	return &Handler{
		Pattern:        "/",
		PatternHandler: &RootHandler{l},
		Method:         "GET",
	}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Println("Attempting access to root site")
	w.WriteHeader(http.StatusNoContent)
}
