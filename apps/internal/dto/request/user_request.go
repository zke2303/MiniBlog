// Package request 请求数据对象
package request

// CreateUserRequest 创建用户请求对象
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=1,max=20,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}
