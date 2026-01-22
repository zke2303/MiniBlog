package controller

import (
	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/dto/request"
	"mini-blog/internal/dto/response"
	"mini-blog/internal/service"

	"github.com/gin-gonic/gin"
)

// BlogController 博文控制器
type BlogController struct {
	svc *service.BlogService
}

// NewBlogController 创建一个 BlogController 控制器
func NewBlogController(svc *service.BlogService) *BlogController {
	return &BlogController{
		svc: svc,
	}
}

// Create 新建一篇博文
func (h *BlogController) Create(c *gin.Context) {
	//  1.绑定数据,并校验
	var req request.CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// 获取当前登入的 userID
	userID := c.GetString("userID")
	if userID == "" {
		c.Error(errmsg.UserNotLogin)
		return
	}
	// 2.调用 service 层
	id, err := h.svc.Create(c.Request.Context(), userID, req)
	if err != nil {
		c.Error(err)
		return
	}

	// 响应 blog id
	response.Success(c, id)
}

// ListBlogs 查询Blogs的概述列表
// @Summary 查询Blogs的概述列表
// @Tags blogs
// @Accept json
// @Produce json
// @Param body {object}
// @Success 200 {object} response.Response{data = }
// @Route /blogs [get]
func (h *BlogController) ListBlogs(c *gin.Context) {
	// 1.绑定并校验请求体参数
	var req request.BlogsPageQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}
	// 2.调用 service 层
	blogs, err := h.svc.ListBlogs(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	// 3.返回查询结果
	response.Success(c, blogs)
}
