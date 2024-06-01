package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"github.com/jonattasmoraes/app-go/internal/utils"
)

func PatchUserHandler(ctx *gin.Context) {
	request := patchUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	utils.ToLowerCaseExcept(&request, "password")

	id := ctx.Param("id")

	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, "param: id is required")
		return
	}

	user := models.User{}

	// Find user
	if err := db.First(&user, id).Error; err != nil {
		logger.Errorf("error getting user: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error getting user")
		return
	}

	// Update user
	if err := db.Model(&user).Updates(request).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	response := CreateUserResponse(user)

	utils.SendSuccess(ctx, "updated user", response)
}
