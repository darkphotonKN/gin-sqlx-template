package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// Importing for side effects - Dont Remove
	// This IS being used!
	_ "github.com/lib/pq"
	"log"
	"os"
)

// NOTE: for global db access, do not remove or move inside a function
var DB *sqlx.DB

func InitDB() *sqlx.DB {

	// construct the db connection string
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("constructed dsn", dsn)

	// pass the db connection string to connect to our database
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Connected to the database successfully.")

	// set global singleton instance for the database
	DB = db
	return DB
}
