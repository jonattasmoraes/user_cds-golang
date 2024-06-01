package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"github.com/jonattasmoraes/app-go/internal/utils"
)

func DeleteUserByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, "param: id is required")
		return
	}

	user := models.User{}

	// Find user
	if err := db.First(&user, id).First(&user).Error; err != nil {
		logger.Errorf("error getting user: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error getting user")
		return
	}

	// Delete user
	if err := db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		logger.Errorf("error deleting user: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting user")
		return
	}

	response := DeleteUserResponse(user)

	utils.SendSuccess(ctx, "user deleted", response)
}
