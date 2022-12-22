package main

import (
	"log"

	"github.com/Aayush9029/PiedPiper/api"
	"github.com/Aayush9029/PiedPiper/api/nodes"
	"github.com/Aayush9029/PiedPiper/database"
	"github.com/Aayush9029/PiedPiper/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Server initialization
	app := server.Create()

	// Migrations
	database.DB.AutoMigrate(&nodes.Node{})

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
