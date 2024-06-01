package config

import (
	"github.com/jonattasmoraes/app-go/internal/database"
	"github.com/jonattasmoraes/app-go/internal/utils"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *utils.Logger
)

func InitializeConfig() error {
	// Initialize Logger
	logger = utils.NewLogger("config")

	var error error
	db, error = database.InitializeDatabase()
	if error != nil {
		logger.Errorf("error connecting to database: %v", error)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *utils.Logger {
	logger = utils.NewLogger(p)
	return logger
}
