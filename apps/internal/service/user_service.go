// Package service 业务处理包
package service

import (
	"context"
	"errors"

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

// GetByID 根据用户id查询用户信息
func (svc *UserService) GetByID(ctx context.Context, userID string) (model.User, error) {
	// 1.校验id格式
	_, err := uuid.Parse(userID)
	if err != nil {
		return model.User{}, errmsg.InternalErr.Wrap(err)
	}

	// 2. 调用 repository 层,执行 sql 查询操作
	user, err := svc.repo.GetByID(ctx, svc.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errmsg.UserNotFound
		}
	}

	// 3.执行成功,返回查询结构
	return user, nil
}
