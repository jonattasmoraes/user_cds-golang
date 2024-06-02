package migrations

import (
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"gorm.io/gorm"
)

func RunMigrations(database *gorm.DB) error {
	// Run migrations
	database.AutoMigrate(&models.User{})
	return nil
}
