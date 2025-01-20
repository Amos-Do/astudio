package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/amosli/astudio/server/author"
	"github.com/amosli/astudio/server/internal/repository/postgres"
	"github.com/amosli/astudio/server/internal/rest"
	"github.com/amosli/astudio/server/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	initEnvSetting()
}

func main() {
	// init logger
	logger := logger.NewLogger()
	defer logger.Close()

	zap.L().Info(
		"set up env config...",
		zap.String("APP", os.Getenv("APP")),
		zap.String("VERSION", os.Getenv("VERSION")),
		zap.String("POSTGRES_HOST", os.Getenv("POSTGRES_HOST")),
	)

	// prepare DB
	db, err := sql.Open(
		os.Getenv("DB_DRV"),
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")),
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

	// prepare gin
	g := gin.Default()

	// prepare repository
	authorRepo := postgres.NewAuthorRepo(db)

	// build service Layer
	authorService := author.NewAuthorService(authorRepo)

	// build rest delivery Layer
	rest.NewAuthorHandler(g, authorService)

	// Start server
	serverAddr := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	zap.S().Infof("Start server %s", serverAddr)
	log.Fatal(g.Run(serverAddr))
}

// initEnvSetting will init environment setting
func initEnvSetting() {
	// export MODE=dev
	//
	// execution specific env with one line
	// $ MODE=dev go run main.go
	env := os.Getenv("MODE")

	// load different mode .env
	// .local > .x > .env
	godotenv.Load(".env." + env + ".local")
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
}
