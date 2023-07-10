package cmshandlers

import (
	"SimpleCMServer/auth"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type request struct {
	GrantType    string `json:"grant_type"`
	Client       string `json:"client"`
	ClientSecret string `json:"client_secret"`
	AuthCode     string `json:"auth_code"`
	RedirectUri  string `json:"redirect_uri"`
}

type response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires"`
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

	if req.GrantType != "authorization_code" {
		w.WriteHeader(http.StatusBadRequest)
	}

	resp := &response{
		AccessToken: "",
		TokenType:   "Bearer",
		Expires:     int(60 * time.Minute),
	}
	// prepare token here

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
