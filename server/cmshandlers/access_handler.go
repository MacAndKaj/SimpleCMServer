package cmshandlers

import (
	"SimpleCMServer/auth"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type request struct {
	grantType    string `json:"grant_type"`
	client       string `json:"client"`
	clientSecret string `json:"client_secret"`
	authCode     string `json:"auth_code"`
	redirectUri  string `json:"redirect_uri"`
}

type response struct {
	accessToken string `json:"access_token"`
	tokenType   string `json:"token_type"`
	expires     int    `json:"expires"`
}

type AccessHandler struct {
	logger      *log.Logger
	authService *auth.AuthService
}

func NewAccessHandler(l *log.Logger) *Handler {
	return &Handler{
		Pattern:        "/oauth/access",
		PatternHandler: &RootHandler{l},
		Method:         "POST",
	}
}

func (h *AccessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.grantType != "authorization_code" {
		w.WriteHeader(http.StatusBadRequest)
	}

	resp := &response{
		accessToken: "",
		tokenType:   "Bearer",
		expires:     int(60 * time.Minute),
	}
	// prepare token here

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
