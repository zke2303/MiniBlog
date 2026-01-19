// Package service 业务处理包
package service

import "mini-blog/internal/repository"

// UserService 用户业务逻辑对象
type UserService struct {
	repo repository.IUserRepository
}

// NewUserService 创建 UserService 实例对象
func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
