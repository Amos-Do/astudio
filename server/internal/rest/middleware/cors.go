package middleware

import (
	"time"

	"github.com/Amos-Do/astudio/server/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS will handle the CORS middleware
func CORS(conf *config.Config) gin.HandlerFunc {
	return cors.New(corsConfig(conf))
}

// corsConfig will set CORS configuration in different env
func corsConfig(conf *config.Config) cors.Config {
	corsConf := cors.Config{
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}

	if conf.Server.Run == gin.DebugMode {
		// develop env
		// allow all settings like origins, methos and headers
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
	} else {
		// production env
		// corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		// corsConf.AllowHeaders = []string{"Authization", "Content-Type", "Origin",
		// "Connection", "Accept-Encoding", "Accept-Language", "Host"}
		// corsConf.AllowOrigins = []string{"https://www.example.com"}
	}
	return corsConf
}
