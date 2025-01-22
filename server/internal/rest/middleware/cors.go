package middleware

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS will handle the CORS middleware
func CORS() gin.HandlerFunc {
	return cors.New(corsConfig())
}

// corsConfig will set CORS configuration in different env
func corsConfig() cors.Config {
	corsConf := cors.Config{
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}

	if os.Getenv("SERVER_RUN_MODE") == "debug" {
		// develop env
		// allow all settings like origins, methos and headers
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
	} else {
		// production env
		// corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		// corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Origin",
		// "Connection", "Accept-Encoding", "Accept-Language", "Host"}
		// corsConf.AllowOrigins = []string{"https://www.example.com"}
	}
	return corsConf
}
