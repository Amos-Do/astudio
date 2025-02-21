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
	g.GET("/ping", handler.Ping)
	g.POST("/auth/login", handler.Login)
	g.POST("/auth/signup", handler.Signup)
	g.GET("/auth/refresh", handler.RefreshToken)
}

// @Summary	Ping check server connection
// @Tags		Auth
// @version	1.0
// @produce	text/plain
// @Success	200	string	string	成功後返回的值
// @Router		/ping [get]
func (h *AuthHandler) Ping(c *gin.Context) {
	res, err := h.Service.Ping(c)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary	Vendor login system
// @Tags		Auth
// @version	1.0
// @produce	application/json
// @param		data	body		domain.LoginRequest	true	"data"
// @Success	200		{object}	domain.AuthToken
// @Router		/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(utils.GetStatusCode(domain.ErrBadParamInput), domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	auth := domain.Auth{
		Account:  req.Email,
		Password: req.Password,
	}
	authToken, err := h.Service.Login(c, auth)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, authToken)
}

// @Summary	Vendor signup system
// @Tags		Auth
// @version	1.0
// @produce	application/json
// @param		email	body		domain.SignupRequest	true	"data"
// @Success	200		{object}	domain.AuthToken
// @Router		/auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req domain.SignupRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(utils.GetStatusCode(domain.ErrBadParamInput), domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	auth := domain.Auth{
		Account:  req.Email,
		Password: req.Password,
	}
	authToken, err := h.Service.Signup(c, auth)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, authToken)
}

// @Summary	Vendor refresh token
// @Tags		Auth
// @version	1.0
// @produce	application/json
// @param		refresh_token	query		string	true	"refresh_token"
// @Success	200				{object}	domain.AuthToken
// @Router		/auth/refresh [get]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req domain.RefreshTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(utils.GetStatusCode(domain.ErrBadParamInput), domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	authToken, err := h.Service.RefreshToken(c, req.RefreshToken)
	if err != nil {
		c.JSON(utils.GetStatusCode(err), domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, authToken)
}
