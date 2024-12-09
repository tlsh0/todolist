package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/tlsh0/todolist/internal/handlers"
)

type app struct {
	routes *http.ServeMux
	cache  *redis.Client
	db     *sql.DB
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	routes := handlers.Routes()

	cache := redis.NewClient(&redis.Options{})

	connString := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("error connecting to the db: ", err)
	}

	http.ListenAndServe(":3000", app.routes)
}
