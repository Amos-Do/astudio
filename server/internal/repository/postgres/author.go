package postgres

import (
	"database/sql"

	"github.com/amosli/astudio/server/domain"
)

type AuthorRepo struct {
	DB *sql.DB
}

// NewAuthorRepo will create an implementation of author.Repository
func NewAuthorRepo(db *sql.DB) domain.IAuthorRepo {
	return &AuthorRepo{DB: db}
}
