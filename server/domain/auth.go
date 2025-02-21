package domain

import (
	"context"
	"database/sql"
	"time"
)

// Auth representing the Auth data struct
type Auth struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Birthday  sql.NullTime `json:"birthday"`
	Account   string       `json:"account"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
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
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}

// IAuthRepo represent the Auth's repository contrat
type IAuthRepo interface {
	GetByEmail(c context.Context, email string) (Auth, error)
	GetByID(c context.Context, id int64) (Auth, error)
	Create(c context.Context, auth *Auth) error
}

// IAuthService represent the Auth's usecases
type IAuthService interface {
	Ping(c context.Context) (string, error)
	Login(c context.Context, auth Auth) (AuthToken, error)
	Signup(c context.Context, auth Auth) (AuthToken, error)
	RefreshToken(c context.Context, refreshToken string) (AuthToken, error)
}
