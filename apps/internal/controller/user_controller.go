// Package controller 控制器
package controller

import (
	"net/http"

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
		response.Fail(c, http.StatusBadRequest, 400, "参数错误")
		return
	}

	id, err := h.svc.Create(c, req)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, 400, "参数错误")
		return
	}

	response.Success(c, id)
}
