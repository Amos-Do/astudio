package author

import (
	"context"

	"github.com/Amos-Do/astudio/server/domain"
)

type Service struct {
	Repo domain.IAuthorRepo
}

// NewAuthorService will create a article service object
func NewAuthorService(repo domain.IAuthorRepo) *Service {
	return &Service{repo}
}

func (s *Service) Ping(c context.Context) (string, error) {
	return "Pong", nil
}
