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
	g.GET("api/v1/author/ping", handler.Ping)
}

// @Summary Ping check server connection
// @Tags Author
// @version 1.0
// @produce text/plain
// @Success 200 string string 成功後返回的值
// @Router /author/ping [get]
func (h *AuthorHandler) Ping(c *gin.Context) {
	res, err := h.Service.Ping(c)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), utils.ResponseErr{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
