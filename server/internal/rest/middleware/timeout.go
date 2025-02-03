package middleware

import (
	"net/http"
	"time"

	"github.com/Amos-Do/astudio/server/internal/rest/utils"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func timeoutResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, utils.ResponseErr{
		Message: "timeout",
	})
}

func SetRequestWithTimeout(d time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(d),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}
