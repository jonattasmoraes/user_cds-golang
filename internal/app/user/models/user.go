package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	LastName string
	Email    string
	Password string
	Role     string
}

type CreateUserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
