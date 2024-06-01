package handler

import "github.com/jonattasmoraes/app-go/internal/app/user/models"

type createUserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateUserResponse(user models.User) createUserResponse {
	return createUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
