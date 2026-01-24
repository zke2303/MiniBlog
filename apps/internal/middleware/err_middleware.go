// Package middleware gin中间件
package middleware

import (
	"encoding/json"
	"fmt"

	"mini-blog/internal/dto/response"
	"mini-blog/internal/pkg/errmsg"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ErrorHandlerMiddleware 错误处理中间件
func ErrorHandlerMiddleware(trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 放行，让 gin 先处理后面的逻辑
		c.Next()
		// 判断是否存在 error
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		// 判断错误类型
		// 1) Validtor 校验错误
		if verrs, ok := err.(validator.ValidationErrors); ok {
			errs := make(map[string]string)
			for _, e := range verrs {
				errs[e.Field()] = e.Translate(trans)
			}

			c.JSON(400, response.Response{
				Code:   errmsg.CodeInvalidParam,
				Msg:    "参数校验错误",
				Errors: errs,
			})
			return
		}

		// 2) JSON 解析错误 (SyntaxError 或 UnmarshalTypeError)
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			response.BizErrFail(c, *errmsg.InvalidParamErr)
			return
		}
		if _, ok := err.(*json.SyntaxError); ok {
			response.BizErrFail(c, *errmsg.InvalidParamErr)
			return
		}

		// 3) 业务错误
		if bizErr, ok := err.(*errmsg.BizErr); ok {
			fmt.Printf("cause error: %v\n", bizErr.Cause)
			response.BizErrFail(c, *bizErr)
			return
		}

		// 3) 未知错误
		response.BizErrFail(c, *errmsg.InternalErr)
	}
}
