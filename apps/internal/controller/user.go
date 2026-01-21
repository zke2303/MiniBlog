// Package controller 控制器
package controller

import (
	"mini-blog/internal/dto/request"
	"mini-blog/internal/dto/response"
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
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	id, err := h.svc.Create(c, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.Success(c, id)
}

// Profile 查询当前登入用户信息
// @Summary 获取当前登录用户的信息
// @Tags users
// @Accept json
// @Produce json
// @Param none
// @Success 200 {object}
// @router /api/v1/auth/profile [get]
func (h *UserController) Profile(c *gin.Context) {
	// 1.获取当前登入的用户id
	// TODO: userID 应该从 context 中获取
	userID := c.GetString("userID")
	// 2.调用 service 层
	user, err := h.svc.GetByID(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}
	// 3.返回用户信息
	response.Success(c, user)
}
