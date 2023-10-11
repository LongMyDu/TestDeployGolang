package api

import (
		"log"
	"os"

	"sample/api/controllers"
	"sample/api/seed"
)

var server = controllers.Server{}

// Run is...
func Run() {
	server.Initalize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"))

	if os.Getenv("ENVIRONMENT") == "DEV" {
		seed.Load(server.DB)
	}

	serverRunning := server.Run(":8080")

	if serverRunning != nil {
		log.Fatalln("Server failed to start")
	}
}
