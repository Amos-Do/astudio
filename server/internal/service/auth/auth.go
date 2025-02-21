package auth

import (
	"context"

	"github.com/Amos-Do/astudio/server/config"
	"github.com/Amos-Do/astudio/server/domain"
	"github.com/Amos-Do/astudio/server/internal/tokenutil"
	"github.com/Amos-Do/astudio/server/pkg/utils"
)

type Service struct {
	Repo domain.IAuthRepo
	Conf *config.Config
}

// NewAuthService will create a article service object
func NewAuthService(conf *config.Config, repo domain.IAuthRepo) domain.IAuthService {
	return &Service{repo, conf}
}

func (s *Service) Ping(c context.Context) (string, error) {
	return "Pong", nil
}

// Login
func (s *Service) Login(c context.Context, auth domain.Auth) (domain.AuthToken, error) {
	// check if the email account not exist
	user, err := s.Repo.GetByEmail(c, auth.Account)
	if err != nil {
		return domain.AuthToken{}, err
	}

	if !utils.CompareHashAndPassword(user.Password, auth.Password) {
		return domain.AuthToken{}, domain.ErrNotAuthized
	}

	accessToken, accessExp, err := tokenutil.CreateAccessToken(&user, s.Conf.Token.AccessSecret, s.Conf.Token.AccessExpiryMs)
	if err != nil {
		return domain.AuthToken{}, err
	}

	refreshToken, refreshExp, err := tokenutil.CreateRefreshToken(&user, s.Conf.Token.RefreshSecret, s.Conf.Token.RefreshExpiryMs)
	if err != nil {
		return domain.AuthToken{}, err
	}

	return domain.AuthToken{
		AccessToken:   accessToken,
		AccessExpiry:  accessExp.Unix(),
		RefreshToken:  refreshToken,
		RefreshExpiry: refreshExp.Unix(),
	}, nil
}

// Signup
func (s *Service) Signup(c context.Context, auth domain.Auth) (domain.AuthToken, error) {
	// check if the email account has been created
	_, err := s.Repo.GetByEmail(c, auth.Account)
	if err != nil {
		if err != domain.ErrNotFound {
			return domain.AuthToken{}, err
		}
	} else {
		return domain.AuthToken{}, domain.ErrConflictCreateExistsAuthAccount
	}

	// encrypted passwordz
	encryptedPassword, err := utils.GenerateFromPassword(auth.Password)
	if err != nil {
		return domain.AuthToken{}, err
	}
	auth.Password = encryptedPassword
	auth.Name = auth.Account // default name

	err = s.Repo.Create(c, &auth)
	if err != nil {
		return domain.AuthToken{}, err
	}

	accessToken, accessExp, err := tokenutil.CreateAccessToken(&auth, s.Conf.Token.AccessSecret, s.Conf.Token.AccessExpiryMs)
	if err != nil {
		return domain.AuthToken{}, err
	}

	refreshToken, refreshExp, err := tokenutil.CreateRefreshToken(&auth, s.Conf.Token.RefreshSecret, s.Conf.Token.RefreshExpiryMs)
	if err != nil {
		return domain.AuthToken{}, err
	}

	return domain.AuthToken{
		AccessToken:   accessToken,
		AccessExpiry:  accessExp.Unix(),
		RefreshToken:  refreshToken,
		RefreshExpiry: refreshExp.Unix(),
	}, nil
}

// RefreshToken
func (s *Service) RefreshToken(c context.Context, refreshToken string) (domain.AuthToken, error) {
	id, err := tokenutil.ExtractIDFromToken(refreshToken, s.Conf.Token.RefreshSecret)
	if err != nil {
		return domain.AuthToken{}, domain.ErrNotAuthized
	}

	auth, err := s.Repo.GetByID(c, id)
	if err != nil {
		return domain.AuthToken{}, err
	}

	accessToken, accessExp, err := tokenutil.CreateAccessToken(&auth, s.Conf.Token.AccessSecret, s.Conf.Token.AccessExpiryMs)
	if err != nil {
		return domain.AuthToken{}, err
	}

	refreshToken, refreshExp, err := tokenutil.CreateRefreshToken(&auth, s.Conf.Token.RefreshSecret, s.Conf.Token.RefreshExpiryMs)
	if err != nil {
		return domain.AuthToken{}, err
	}

	return domain.AuthToken{
		AccessToken:   accessToken,
		AccessExpiry:  accessExp.Unix(),
		RefreshToken:  refreshToken,
		RefreshExpiry: refreshExp.Unix(),
	}, nil
}
