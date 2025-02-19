package utils

import (
	"net/http"

	"github.com/Amos-Do/astudio/server/domain"
	"go.uber.org/zap"
)

// GetStatusCode will map domain errors to http status codes
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	zap.S().Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict,
		domain.ErrConflictCreateExistsAuthAccount:
		return http.StatusConflict
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	case domain.ErrNotAuthized:
		return http.StatusUnauthorized
	case domain.ErrGetwayTimeout:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}
