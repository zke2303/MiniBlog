// Package controller 控制器
package controller

import (
	"mini-blog/internal/service"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	svc *service.UserService
}

// NewUserController 创建 UserController 实例
func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

// Create 创建用户实例对象
func (h *UserController) Create(c *gin.Context) {
}
