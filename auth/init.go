package auth

import (
	"database/sql"
	"log"
)

type AuthService struct {
	logger  *log.Logger
	usersDb *sql.DB
}

func NewAuthService(l *log.Logger, d *sql.DB) *AuthService {
	err := d.Ping()
	if err != nil {
		l.Fatal("Cannot ping database!")
		return nil
	}

	service := &AuthService{
		logger:  l,
		usersDb: d,
	}

	return service
}
