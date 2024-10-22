package main

import (
	"os"

	"github.com/darkphotonKN/gin-sqlx-template/config"
)

// main entry point to app
func main() {

	// db setup
	db := config.InitDB()
	defer db.Close()

	// router setup
	router := config.SetupRouter()

	defaultDevPort := ":8080"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultDevPort
	}

	// run server listener
	router.Run(port)
}
