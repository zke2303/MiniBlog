package request

// CreateBlogRequest 创建 Blog 请求体
type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=30"`
	Content string `json:"content" binding:"required,min=20"`
}
