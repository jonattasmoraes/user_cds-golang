package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %v", err)
		return nil, err
	}

	// String to connect to database
	dsn := os.Getenv("DSN")

	// Connect to database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("error connecting to database: %v", err)
		return nil, err
	}

	// Return database
	return database, err
}
