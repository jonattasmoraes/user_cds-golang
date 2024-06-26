package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"github.com/jonattasmoraes/app-go/internal/utils"
)

// @BasePath /api

// @Summary Get All Users
// @Description List all users
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} GetUsersResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/ [get]
func GetAllUsersHandler(ctx *gin.Context) {
	users := []models.User{}

	if err := db.Find(&users).Error; err != nil {
		logger.Errorf("error getting users: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error getting users")
		return
	}

	response := GetAllUsersResponse(users)

	utils.SendSuccess(ctx, "list of users", response)
}
