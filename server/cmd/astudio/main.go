package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	initEnvSetting()
}

func main() {

	// 連接 DB
	db, err := sql.Open(
		os.Getenv("DB_DRV"),
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")),
	)
	if err != nil {
		panic(err)
	}

	// check db connect
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully created connection to database")

	// run a simple http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!")
	})

	serverAddr := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	fmt.Println("Start server", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

// init environment setting
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
