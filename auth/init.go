package auth

import (
	"database/sql"
	"log"
)

type AuthService struct {
	logger   *log.Logger
	tokensDb *sql.DB
}

func NewAuthService(l *log.Logger, d *sql.DB) *AuthService {
	err := d.Ping()
	if err != nil {
		l.Fatal("Cannot ping database!")
		return nil
	}

	service := &AuthService{
		logger:   l,
		tokensDb: d,
	}

	return service
}
