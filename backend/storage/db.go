package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	DB, err = sql.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		return err
	}
	return nil
}