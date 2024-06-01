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
