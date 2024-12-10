package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/tlsh0/todolist/internal/handlers"
	repo "github.com/tlsh0/todolist/internal/repository"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// config
	// routes
	// caching
	//
	//
	routes := handlers.Routes()

	// cache := redis.NewClient(&redis.Options{})

	connString := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to the db: ", err)
	}

	storage := repo.New(db)

	storage.CreateUser()

	http.ListenAndServe(":3000", routes)
}
