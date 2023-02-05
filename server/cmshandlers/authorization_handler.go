package cmshandlers

import (
	"log"
	"net/http"
)

type AuthorizationHandler struct {
	logger *log.Logger
}

func NewAuthorizationHandler(l *log.Logger) *Handler {
	return &Handler{
		Pattern:        "/oauth/authorize",
		PatternHandler: &RootHandler{l},
		Method:         "GET",
	}
}

func (h *AuthorizationHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {

}
