package auth

import (
	"context"

	"github.com/Amos-Do/astudio/server/domain"
)

type Service struct {
	Repo domain.IAuthRepo
}

// NewAuthService will create a article service object
func NewAuthService(repo domain.IAuthRepo) *Service {
	return &Service{repo}
}

func (s *Service) Ping(c context.Context) (string, error) {
	return "Pong", nil
}
