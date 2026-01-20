// Package response 请求响应对象
package response

import (
	"net/http"

	"mini-blog/internal/dto/errmsg"

	"github.com/gin-gonic/gin"
)

// Response 请求响应对象
type Response struct {
	Code   int    `json:"code"`             // 错误码
	Msg    string `json:"msg"`              // 错误描述
	Data   any    `json:"data,omitempty"`   // 响应数据
	Errors any    `json:"errors,omitempty"` // 错误详细描述
}

// New 快速创建 Response 对象
func New(code int, msg string, data any) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// Success 响应成功
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// Fail 响应失败
func Fail(c *gin.Context, statusCode int, code int, msg string) {
	c.JSON(statusCode, Response{
		Code: code,
		Msg:  msg,
	})
}

// BizErrFail 业务异常
func BizErrFail(c *gin.Context, bizErr errmsg.BizErr) {
	c.JSON(http.StatusOK, Response{
		Code: bizErr.Code,
		Msg:  bizErr.Msg,
	})
}
