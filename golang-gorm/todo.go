package golang_gorm

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          string         `gorm:"primary_key;column:id;autoIncrement"`
	UserId      string         `gorm:"column:user_id"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (todo *Todo) TableName() string {
	return "todos"
}
