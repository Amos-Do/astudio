package middleware

import (
	"strings"

	"github.com/Amos-Do/astudio/server/domain"
	"github.com/Amos-Do/astudio/server/internal/rest/utils"
	"github.com/Amos-Do/astudio/server/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

// JwtAuth will handle Authorization jwt token middleware
func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		// check Authorization param
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)

			// check the token is authorized
			if authorized {
				// extract the ID from token
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(utils.GetStatusCode(domain.ErrNotAuthorized), domain.ErrorResponse{
						Message: err.Error(),
					})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(utils.GetStatusCode(domain.ErrNotAuthorized), domain.ErrorResponse{
				Message: err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(utils.GetStatusCode(domain.ErrNotAuthorized), domain.ErrorResponse{
			Message: domain.ErrNotAuthorized.Error(),
		})
		c.Abort()
	}
}
