package domain

import "context"

// Author representing the Author data struct
type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// IAuthorRepo represent the Author's repository contrat
type IAuthorRepo interface {
}

// IAuthorService represent the Author's usecases
type IAuthorService interface {
	Ping(c context.Context) (string, error)
}
