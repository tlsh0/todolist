package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tlsh0/todolist/internal/db"
	"github.com/tlsh0/todolist/internal/models"
	"github.com/tlsh0/todolist/internal/router"
)

func main() {
	// getting the flags for the migration
	tablesUp := flag.Bool("up", false, "Use 'up' to get the tables up in the DB")
	tablesDown := flag.Bool("down", false, "Use 'down' to get the tables down in the DB")

	flag.Parse()

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// getting up the tables
	if *tablesUp {
		err = db.AutoMigrate(&models.User{}, &models.Task{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("tables up successfully!")
	}

	// dropping all the tables
	if *tablesDown {
		err = db.Migrator().DropTable(&models.User{}, &models.Task{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("tables down...")
	}

	r := router.NewRouter()
	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
