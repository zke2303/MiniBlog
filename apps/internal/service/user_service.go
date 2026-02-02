// Package service 业务处理包
package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"mini-blog/internal/dto/request"
	"mini-blog/internal/model"
	"mini-blog/internal/pkg/errmsg"
	"mini-blog/internal/pkg/utils"
	"mini-blog/internal/repository"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户业务逻辑对象
type UserService struct {
	db         *gorm.DB
	rdb        *redis.Client
	repo       repository.IUserRepository
	jwtHandler *utils.JwtHandler
}

// NewUserService 创建 UserService 实例对象
func NewUserService(db *gorm.DB,
	rdb *redis.Client,
	repo repository.IUserRepository,
	jwtHandler *utils.JwtHandler,
) *UserService {
	return &UserService{
		db:         db,
		rdb:        rdb,
		repo:       repo,
		jwtHandler: jwtHandler,
	}
}

// Create 创建用户
func (svc *UserService) Create(ctx context.Context, req request.CreateUserRequest) (string, error) {
	// 1. 生成 ID
	uid, err := uuid.NewV7()
	if err != nil {
		return "", errmsg.New(errmsg.CodeInternalErr, "生成用户 ID 失败", err)
	}

	// 2. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errmsg.New(errmsg.CodeInternalErr, "密码加密失败", err)
	}

	// 3. 构建模型
	user := model.User{
		ID:       uid,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	// 4. 执行业务逻辑
	err = svc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := svc.repo.GetByUsername(ctx, tx, req.Username)
		if err == nil {
			return errmsg.UserAlreadyExists
		}
		if !errors.Is(err, errmsg.UserNotFound) && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := svc.repo.Create(ctx, tx, user); err != nil {
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "23505") {
				return errmsg.UserAlreadyExists
			}
			return err
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

// GetByID 根据用户id查询用户信息 (带缓存优化)
func (svc *UserService) GetByID(ctx context.Context, userID string) (model.User, error) {
	cacheKey := fmt.Sprintf("user:profile:%s", userID)

	// 1. 尝试从 Redis 获取
	val, err := svc.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var user model.User
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return user, nil
		}
	}

	// 2. 校验id格式
	if _, err := uuid.Parse(userID); err != nil {
		return model.User{}, errmsg.InternalErr.Wrap(err)
	}

	// 3. 调用 repository 层
	user, err := svc.repo.GetByID(ctx, svc.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errmsg.UserNotFound
		}
		return model.User{}, err
	}

	// 4. 写入缓存 (过期时间 10 分钟)
	userJson, _ := json.Marshal(user)
	svc.rdb.Set(ctx, cacheKey, userJson, 10*time.Minute)

	return user, nil
}

// Login 用户登入
func (svc *UserService) Login(ctx context.Context, req request.UserLoginRequest) (string, error) {
	user, err := svc.repo.GetByUsername(ctx, svc.db, req.Username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errmsg.New(errmsg.CodePasswordError, "密码错误", err)
	}

	token, err := svc.jwtHandler.GeneratorJwt(user.ID.String(), user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Update 更新用户信息
func (svc *UserService) Update(ctx context.Context, userID string, req request.UpdateUserRequest) error {
	if req.Password != nil {
		password, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return errmsg.InternalErr.Wrap(err)
		}
		passwordStr := string(password)
		req.Password = &passwordStr
	}

	err := svc.db.Transaction(func(db *gorm.DB) error {
		if err := svc.repo.Update(ctx, db, userID, req); err != nil {
			return err
		}
		// 更新后清除缓存
		cacheKey := fmt.Sprintf("user:profile:%s", userID)
		svc.rdb.Del(ctx, cacheKey)
		return nil
	})

	return err
}
