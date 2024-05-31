package config

import (
	"fmt"

	"github.com/jonattasmoraes/app-go/internal/database"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeConfig() error {
	var error error
	db, error = database.InitializeDatabase()
	if error != nil {
		fmt.Printf("error connecting to database: %v", error)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
