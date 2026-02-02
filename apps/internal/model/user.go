package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"create_time" gorm:"column:create_time"`
	UpdatedAt time.Time      `json:"update_time" gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"column:delete_time;index"`
}
