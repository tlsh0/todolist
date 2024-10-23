package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBcredentials struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

func Connect() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// reading the credentials from .env
	dbCredentials := &DBcredentials{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	// connecting to the DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbCredentials.Host, dbCredentials.User, dbCredentials.Password, dbCredentials.Name, dbCredentials.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
