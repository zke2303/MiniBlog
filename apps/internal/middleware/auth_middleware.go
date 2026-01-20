package middleware

import (
	"net/http"
	"strings"

	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从请求中获取 token
		header := c.Request.Header.Get("Authorization")
		// 2.判断是否存在请求头
		if header == "" {
			// 没有携带请求头
			_ = c.Error(errmsg.New(http.StatusBadRequest, "未携带请求头", nil))
			c.Abort()
			return
		}

		// 3.判断请求头格式是否正确
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			_ = c.Error(errmsg.New(http.StatusBadRequest, "请求头格式错误", nil))
			c.Abort()
			return
		}

		// 4.解析 token
		claims, err := utils.PasrseJwt(parts[1])
		if err != nil {
			_ = c.Error(errmsg.New(http.StatusUnauthorized, "token解析失败", nil))
			c.Abort()
			return
		}

		// 5.将关键信息存入 c.Context中
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
