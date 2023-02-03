package cmshandlers

import (
	"log"
	"net/http"
)

type Handler struct {
	Pattern        string
	PatternHandler http.Handler
}

func BuildHandlers(l *log.Logger) []*Handler {
	handlers := []*Handler{
		NewRootHandler(l),
	}
	return handlers
}
