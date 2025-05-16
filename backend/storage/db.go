package storage

import (
	"database/sql"
	"fmt"
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

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPass)

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		return err
	}
	return nil
}