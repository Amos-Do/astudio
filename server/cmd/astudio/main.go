package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Amos-Do/astudio/server/config"
	"github.com/Amos-Do/astudio/server/docs"

	"github.com/Amos-Do/astudio/server/internal/repository/postgres"
	"github.com/Amos-Do/astudio/server/internal/rest/middleware"
	"github.com/Amos-Do/astudio/server/internal/rest/v1"
	"github.com/Amos-Do/astudio/server/internal/service/auth"
	"github.com/Amos-Do/astudio/server/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title           a_studio API
// @description     This is a a_studio server celler server.

// @contact.name   Amos Li
// @contact.url    https://amos-do.github.io/AmosLi/index.html
// @contact.email  amosli.sj@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
// schemes http
func main() {
	conf := config.New()

	// init logger
	logger := logger.New(
		logger.SetLogLevel(conf.Log.Level),
		logger.SetLogSavePath(conf.Log.SavePath),
		logger.SetLogFileName(conf.Log.FileName),
		logger.SetLogFileExt(conf.Log.FileExt),
	)
	defer logger.Close()

	zap.L().Info(
		"set up env config...",
		zap.String("APP", conf.App.Name),
		zap.String("VERSION", conf.App.Version),
		zap.String("DB_HOST", conf.DB.Host),
	)

	// prepare DB
	db, err := sql.Open(
		conf.DB.Type,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Password, conf.DB.Name),
	)
	if err != nil {
		zap.L().Error("failed to connect db", zap.Error(err))
		panic(err)
	}

	// check db connect
	if err = db.Ping(); err != nil {
		zap.L().Error("failed to check db connect", zap.Error(err))
		panic(err)
	}
	zap.L().Info("Successfully created connection to database")

	// prepare repository
	authRepo := postgres.NewAuthRepo(db)

	// build service Layer
	authService := auth.NewAuthService(authRepo)

	// prepare gin
	g := gin.Default()
	g.Use(middleware.CORS(conf))
	g.Use(middleware.SetRequestWithTimeout(3000 * time.Millisecond))

	// build rest delivery Layer
	// register v1 routes
	rest.SetupV1Api(conf, g, rest.Usecase{
		AuthService: authService,
	})

	// prepare swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", conf.Server.Port)
	docs.SwaggerInfo.Version = conf.App.Version

	// Start server
	serverAddr := fmt.Sprintf(":%s", conf.Server.Port)
	zap.S().Infof("Start server %s", serverAddr)
	log.Fatal(g.Run(serverAddr))
}
