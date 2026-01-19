// Package repository 数据库操作
package repository

import (
	"context"

	"mini-blog/internal/model"

	"gorm.io/gorm"
)

// IUserRepository UserRepository接口
type IUserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user model.User) error
}

// UserRepository UserRepository 实例化对象
type UserRepository struct{}

// NewUserRepository 创建 UserRepository 方法
func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

// Create 创建一条 User 记录
func (repo *UserRepository) Create(ctx context.Context, db *gorm.DB, user model.User) error {
	result := db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
