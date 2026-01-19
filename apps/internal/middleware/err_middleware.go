// Package middleware gin中间件
package middleware

import (
	"errors"
	"fmt"

	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/dto/response"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware 错误处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 放行，让 gin 先处理后面的逻辑
		c.Next()
		// 判断是否存在 error
		if len(c.Errors) > 0 {
			// 获取 gin 调用链中的最后一个错误
			last := c.Errors.Last()
			// 判断错误类型
			var bizErr *errmsg.BizErr
			// 如果是 自定义错误， 返回 自定义错误的信息
			if errors.As(last, &bizErr) {
				response.BizErrFail(c, *bizErr)
			} else {
				// 如果是系统级错误（如数据库崩溃、空指针），屏蔽细节，返回“服务器内部错误”
				// 同时可以在这里打印日志方便排查
				fmt.Printf("[Internal Error] %v\n", last.Err)
				response.BizErrFail(c, errmsg.InternalErr)
			}
			c.Abort()
		}
	}
}
