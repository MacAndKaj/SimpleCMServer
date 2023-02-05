package cmshandlers

import (
	"log"
	"net/http"
)

type Handler struct {
	Pattern        string
	PatternHandler http.Handler
	Method         string
}

func BuildHandlers(l *log.Logger) []*Handler {
	handlers := []*Handler{
		NewRootHandler(l),
		NewPingHandler(l),
		NewAccessHandler(l),
		NewAuthorizationHandler(l),
	}
	return handlers
}
