package rest

import (
	"github.com/Amos-Do/astudio/server/config"
	"github.com/Amos-Do/astudio/server/domain"
	"github.com/Amos-Do/astudio/server/internal/rest/middleware"
	"github.com/gin-gonic/gin"
)

type Usecase struct {
	AuthorService domain.IAuthorService
}

// SetupV1Api handle the all 'v1' api router initialization
func SetupV1Api(conf *config.Config, g *gin.Engine, usecase Usecase) {
	v1 := g.Group("/api/v1")

	// Public APIs
	NewAuthorV1Handler(v1, usecase.AuthorService)

	// Protect APIs
	protectRouter := v1.Group("")

	// middleware to verify AccessToken
	protectRouter.Use(middleware.JwtAuth(conf.Token.AccessSecret))

}
