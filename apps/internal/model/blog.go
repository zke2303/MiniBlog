// Package model 对应于数据库中的字段
package model

import (
	"time"

	"github.com/google/uuid"
)

// Blog 博文对象
type Blog struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_time" gorm:"create_time"`
	UpdateAt time.Time `json:"update_time" gorm:"update_time"`
	DeleteAt time.Time `json:"delete_time" gorm:"delete_time"`
}

// TableName 设置表名
func (*Blog) TableName() string {
	return "blogs"
}
