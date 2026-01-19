// Package service 业务处理包
package service

import (
	"context"

	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/dto/request"
	"mini-blog/internal/model"
	"mini-blog/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户业务逻辑对象
type UserService struct {
	db   *gorm.DB
	repo repository.IUserRepository
}

// NewUserService 创建 UserService 实例对象
func NewUserService(db *gorm.DB, repo repository.IUserRepository) *UserService {
	return &UserService{
		db:   db,
		repo: repo,
	}
}

// Create 创建用户
func (svc *UserService) Create(ctx context.Context, req request.CreateUserRequest) (string, error) {
	// 2.生成 uuid
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", errmsg.New(errmsg.CodeInternal, "创建 uuid 错误", err)
	}

	// 3.加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errmsg.New(errmsg.CodeInternal, "加密密码错误", err)
	}
	// 4.构建 User 对象
	user := model.User{
		ID:       uuid,
		Username: req.Username,
		Password: string(password),
	}

	// 5.调用repository层, 使用事务
	err = svc.db.Transaction(func(tx *gorm.DB) error {
		if err := svc.repo.Create(ctx, svc.db, user); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", errmsg.InternalErr.Wrap(err)
	}

	return uuid.String(), nil
}
