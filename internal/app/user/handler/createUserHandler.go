package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"github.com/jonattasmoraes/app-go/internal/utils"
)

func CreateUserHandler(ctx *gin.Context) {
	request := createUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	utils.ToLowerCaseExcept(&request, "password")

	user := models.User{
		Name:     request.Name,
		LastName: request.LastName,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error creating user")
		return
	}

	response := CreateUserResponse(user)

	utils.SendSuccess(ctx, "user", response)
}
