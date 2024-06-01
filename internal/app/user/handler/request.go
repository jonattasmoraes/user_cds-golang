package handler

import (
	"fmt"

	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"github.com/jonattasmoraes/app-go/internal/utils"
)

type createUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (r *createUserRequest) Validate() error {
	if r.Name == "" && r.LastName == "" && r.Email == "" && r.Password == "" && r.Role == "" {
		return fmt.Errorf("request body is empty or invalid, please try again")
	}

	if r.Name == "" {
		return utils.ParamIsRequired("name", "string")
	}

	if r.LastName == "" {
		return utils.ParamIsRequired("lastname", "string")
	}

	if r.Password == "" {
		return utils.ParamIsRequired("password", "string")
	}

	if r.Role != "admin" && r.Role != "user" && r.Role != "supervisor" {
		return fmt.Errorf("param: role must be 'admin', 'supervisor' or 'user'")
	}

	if r.Email == "" {
		return utils.ParamIsRequired("email", "string")
	}

	if !utils.IsValidEmail(r.Email) {
		return fmt.Errorf("invalid email, please enter a valid email and try again")
	}

	if db.Where("email = ?", r.Email).First(&models.User{}).Error == nil {
		return fmt.Errorf("a user with email '%s' already exists", r.Email)
	}

	return nil
}
