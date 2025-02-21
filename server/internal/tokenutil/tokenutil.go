package tokenutil

import (
	"fmt"
	"time"

	"github.com/Amos-Do/astudio/server/domain"
	"github.com/golang-jwt/jwt/v5"
)

// CreateAccessToken will create jwt token with custom claims
func CreateAccessToken(user *domain.Auth, secret string, expiry int) (accessToken string, exp time.Time, err error) {
	exp = time.Now().Add(time.Millisecond * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", exp, err
	}
	return accessToken, exp, nil
}

// CreateRefreshToken will create jwt token with custom claims
func CreateRefreshToken(user *domain.Auth, secret string, expriy int) (refreshToken string, exp time.Time, err error) {
	exp = time.Now().Add(time.Millisecond * time.Duration(expriy))
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", exp, err
	}
	return refreshToken, exp, nil
}

// IsAuthized will parse the token to validate the signature with token secret
func IsAuthized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// ExtractIDFromToken will extract the ID from token
func ExtractIDFromToken(requestToken string, secret string) (int64, error) {
	token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return int64(claims["id"].(float64)), nil
}
