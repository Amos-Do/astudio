package domain

import "context"

// Auth representing the Auth data struct
type Auth struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// IAuthRepo represent the Auth's repository contrat
type IAuthRepo interface {
}

// IAuthService represent the Auth's usecases
type IAuthService interface {
	Ping(c context.Context) (string, error)
}
