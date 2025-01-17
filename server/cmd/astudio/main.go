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

const (
	HOST     = "postgres_db"
	DATABASE = "a_studio"
	USER     = "admin"
	PASSWORD = "admin"
)

func main() {
	// 連接 DB
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE),
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
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// init environment setting
func initEnvSetting() {
	// export MODE=dev
	env := os.Getenv("MODE")
	if env == "" {
		// default env
		env = "dev"
	}

	// load different mode .env
	// .local > .x > .env
	godotenv.Load(".env." + env + ".local")
	if env != "beta" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env

	fmt.Println("APP: ", os.Getenv("APP"))
	fmt.Println("VERSION: ", os.Getenv("VERSION"))
	fmt.Println("DB: ", os.Getenv("DB"))
}
