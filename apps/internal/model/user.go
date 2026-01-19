package model

import (
	"time"

	"github.com/google/uuid"
)

// User 用户模型
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_time" gorm:"create_time"`
	UpdateAt time.Time `json:"update_time" gorm:"update_time"`
	DeleteAt time.Time `json:"delete_time" gorm:"delete_time"`
}
