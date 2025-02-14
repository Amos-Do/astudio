package rest

import (
	"github.com/Amos-Do/astudio/server/config"
	"github.com/Amos-Do/astudio/server/domain"
	"github.com/Amos-Do/astudio/server/internal/rest/middleware"
	"github.com/gin-gonic/gin"
)

type Usecase struct {
	AuthService domain.IAuthService
}

// SetupV1Api handle the all 'v1' api router initialization
func SetupV1Api(conf *config.Config, g *gin.Engine, usecase Usecase) {
	v1 := g.Group("/api/v1")

	// Public APIs
	NewAuthV1Handler(v1, usecase.AuthService)

	// Protect APIs
	protectRouter := v1.Group("")

	// middleware to verify AccessToken
	protectRouter.Use(middleware.JwtAuth(conf.Token.AccessSecret))

}
