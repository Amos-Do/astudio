package rest

import (
	"net/http"

	"github.com/Amos-Do/astudio/server/domain"
	"github.com/Amos-Do/astudio/server/internal/rest/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service domain.IAuthService
}

// NewAuthV1Handler will initialize the auth/ resources endpoint
func NewAuthV1Handler(g *gin.RouterGroup, svc domain.IAuthService) {
	handler := &AuthHandler{
		Service: svc,
	}
	g.GET("/auth/ping", handler.Ping)
}

// @Summary Ping check server connection
// @Tags Auth
// @version 1.0
// @produce text/plain
// @Success 200 string string 成功後返回的值
// @Router /auth/ping [get]
func (h *AuthHandler) Ping(c *gin.Context) {
	res, err := h.Service.Ping(c)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
