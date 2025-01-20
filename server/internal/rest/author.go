package rest

import (
	"net/http"

	"github.com/amosli/astudio/server/domain"
	"github.com/amosli/astudio/server/internal/rest/utils"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	Service domain.IAuthorService
}

// NewAuthorHandler will initialize the author/ resources endpoint
func NewAuthorHandler(g *gin.Engine, svc domain.IAuthorService) {
	handler := &AuthorHandler{
		Service: svc,
	}
	g.GET("/author/ping", handler.Ping)
}

// Ping will check server connection
func (h *AuthorHandler) Ping(c *gin.Context) {
	res, err := h.Service.Ping(c)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), utils.ResponseErr{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
