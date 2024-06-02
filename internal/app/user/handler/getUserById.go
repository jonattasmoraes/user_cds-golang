package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"github.com/jonattasmoraes/app-go/internal/utils"
)

// @BasePath /api

// @Summary Get User by ID
// @Description Get a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		logger.Errorf("param: id is required")
		utils.SendError(ctx, http.StatusBadRequest, "param: id is required")
		return
	}

	user := models.User{}

	// Find user
	if err := db.First(&user, id).Error; err != nil {
		logger.Errorf("error getting user: %v", err.Error())
		utils.SendError(ctx, http.StatusNotFound, "error getting user")
		return
	}

	response := CreateUserResponse(user)

	utils.SendSuccess(ctx, "get user by id", response)
}
