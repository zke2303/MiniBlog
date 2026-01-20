package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"create_time" gorm:"column:create_time"`
	UpdatedAt time.Time      `json:"update_time" gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"column:delete_time"`
}
