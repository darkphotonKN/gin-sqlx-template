package config

import (
	"github.com/jmoiron/sqlx"
	"log"
)

/**
* Sets up all migration steps for all tables.
**/
func RunMigrations(db *sqlx.DB) {
	query := `
	 CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations ran successfully.")
}
