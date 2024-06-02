package handler

import "github.com/jonattasmoraes/app-go/internal/app/user/models"

func CreateUserResponse(user models.User) models.CreateUserResponse {
	return models.CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func GetAllUsersResponse(users []models.User) []models.CreateUserResponse {
	var response []models.CreateUserResponse
	for _, user := range users {
		response = append(response, CreateUserResponse(user))
	}
	return response
}

func DeleteUserResponse(user models.User) models.CreateUserResponse {
	return models.CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type UserResponse struct {
	Message string                    `json:"message"`
	Data    models.CreateUserResponse `json:"data"`
}

type GetUsersResponse struct {
	Message string                      `json:"message"`
	Data    []models.CreateUserResponse `json:"data"`
}
