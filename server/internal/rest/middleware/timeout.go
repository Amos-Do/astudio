package middleware

import (
	"time"

	"github.com/Amos-Do/astudio/server/domain"
	"github.com/Amos-Do/astudio/server/internal/rest/utils"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// SetRequestWithTimeout will handle the request with timeout
func SetRequestWithTimeout(d time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(d),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.JSON(utils.GetStatusCode(domain.ErrGetwayTimeout), domain.ErrorResponse{
				Message: domain.ErrGetwayTimeout.Error(),
			})
		}),
	)
}
