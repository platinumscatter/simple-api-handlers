package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	isDone bool   `json:"is_done"`
}
