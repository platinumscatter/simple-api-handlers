package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
}