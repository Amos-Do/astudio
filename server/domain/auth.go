package domain

import "context"

// Auth representing the Auth data struct
type Auth struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthToken struct {
	AccessToken   string `json:"access_token"`
	AccessExpiry  int64  `json:"access_expiry"`
	RefreshToken  string `json:"refresh_token"`
	RefreshExpiry int64  `json:"refresh_expiry"`
}

// Auth represent the Auth's delivery request and response
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// IAuthRepo represent the Auth's repository contrat
type IAuthRepo interface {
}

// IAuthService represent the Auth's usecases
type IAuthService interface {
	Ping(c context.Context) (string, error)
}
