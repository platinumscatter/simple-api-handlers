package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Tasks    []Task `json:"tasks"`
}