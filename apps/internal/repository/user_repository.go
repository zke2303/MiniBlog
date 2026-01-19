// Package repository 数据库操作
package repository

import (
	"context"
)

// IUserRepository UserRepository接口
type IUserRepository interface{}

// UserRepository UserRepository 实例化对象
type UserRepository struct{}

// NewUserRepository 创建 UserRepository 方法
func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

// Create 创建一条 User 记录
func (repo *UserRepository) Create(ctx context.Context) (string, error) {
}
