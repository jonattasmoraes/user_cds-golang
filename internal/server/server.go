package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeRouter() {
	// Instantiate server
	server := gin.Default()

	// Initialize routes
	initializeRoutes(server)

	// Get server port
	port := getServerPort()

	//Run server
	server.Run(":" + port)
}

// Get server port from environment
func getServerPort() string {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port " + port)

	return port
}
