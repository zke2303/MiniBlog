package request

// UserLoginRequest 用户登入请求体
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}
