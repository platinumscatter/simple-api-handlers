package taskService

import (
	"github.com/platinumscatter/simple_api/internal/models"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string      `json:"task"`
	IsDone bool        `json:"is_done"`
	UserID uint        `json:"user_id"`
	User   models.User `json:"user"`
}
