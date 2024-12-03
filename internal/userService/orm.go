package userService

import (
	"github.com/platinumscatter/simple_api/internal/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Tasks    []models.Task `json:"tasks"`
}
