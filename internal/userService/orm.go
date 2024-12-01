package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"user"`
	Password string `json:"password"`
}
