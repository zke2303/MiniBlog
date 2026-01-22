// Package model 对应于数据库中的字段
package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Blog 博文对象
type Blog struct {
	ID        uuid.UUID      `json:"id"`
	Title     string         `json:"title"`
	AuthorID  uuid.UUID      `json:"author_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"create_time" gorm:"column:create_time"`
	UpdatedAt time.Time      `json:"update_time" gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt `json:"delete_time" gorm:"column:delete_time"`
}

// TableName 设置表名
func (*Blog) TableName() string {
	return "blogs"
}
