package main

import (
	"db"
	"handlers"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// If a .env file exist, load it
	prod := os.Getenv("PROD")

	if prod != "true" {
		// load .env file
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbType := os.Getenv("DB_TYPE")

	db.SetDataBaseConnector(dbType)

	if !db.DbConn.IsConnection() {
		log.Fatalln("No connection to the DB")
	}

	log.Println("Connection successful to the DB")

	handlers.Handlers()
}
