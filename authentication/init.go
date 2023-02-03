package authentication

import (
	"database/sql"
	"log"
)

type AuthService struct {
	logger *log.Logger
	db     *sql.DB
}

func NewAuthService() *AuthService {
	return nil
}
