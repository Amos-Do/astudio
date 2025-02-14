package postgres

import (
	"database/sql"

	"github.com/Amos-Do/astudio/server/domain"
)

type AuthRepo struct {
	DB *sql.DB
}

// NewAuthRepo will create an implementation of auth.Repository
func NewAuthRepo(db *sql.DB) domain.IAuthRepo {
	return &AuthRepo{DB: db}
}
