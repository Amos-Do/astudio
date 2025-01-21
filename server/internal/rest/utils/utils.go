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
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
