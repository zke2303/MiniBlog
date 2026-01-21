package controller

import (
	"mini-blog/internal/dto/request"
	"mini-blog/internal/dto/response"
	"mini-blog/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthController 鉴权相关 Controller
type AuthController struct {
	svc *service.UserService
}

// NewAuthController 创建 AuthController 实例对象
func NewAuthController(svc *service.UserService) *AuthController {
	return &AuthController{
		svc: svc,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Tags auth
// @Accept json
// @Pruduce json
// @Param body {object} request.CreateUserRequest true
// @Success 200 {object} response.Response{data = string}
// @Router /auth/register
func (h *AuthController) Register(c *gin.Context) {
	// 1.绑定参数，并校验
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// 2. 调用 service 层
	userID, err := h.svc.Create(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	// 3.返回成功信息
	response.Success(c, userID)
}

// Login 用户登入
// @Summary 用户登入
// @Tags auth
// @Accept json
// @Produce json
// @Param
// @Success 200 {object} response.Response{data: string} "登入成功"
// @Router /auth/login
func (h *AuthController) Login(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	token, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, token)
}
