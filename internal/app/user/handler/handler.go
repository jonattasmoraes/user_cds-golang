package handler

import (
	"github.com/jonattasmoraes/app-go/internal/config"
	"github.com/jonattasmoraes/app-go/internal/utils"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *utils.Logger
)

func InitializeHandler() {
	db = config.GetDB()
	logger = utils.NewLogger("Handler")
}
