// Package errmsg 统一错误码
package errmsg

// BizErr 业务错误
type BizErr struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Cause error  `json:"cause"`
}

// Error 实现 Error 接口
func (bizErr *BizErr) Error() string {
	return bizErr.Msg
}

// Wrap 保持异常
func (bizErr BizErr) Wrap(err error) error {
	return &BizErr{
		Code:  bizErr.Code,
		Msg:   bizErr.Msg,
		Cause: err,
	}
}

// New 创建一个新的 错误 实例对象
func New(code int, msg string, cause error) error {
	return &BizErr{
		Code:  code,
		Msg:   msg,
		Cause: cause,
	}
}

// NewInternalErr 快速创建一个带消息的内部服务错误
func NewInternalErr(msg string, cause error) error {
	return &BizErr{
		Code:  CodeInternalErr,
		Msg:   msg,
		Cause: cause,
	}
}

// 定义错误码
const (
	CodeInternalErr      = 10000
	CodeInvalidParam     = 10001
	CodeUUIDGeneratorErr = 10002
	CodeNotLogin         = 10003
	CodeUnmarshalErr     = 10004

	CodeUserNotFoune     = 20001
	CodePasswordError    = 20002
	CodeUserAleadyExists = 20003

	CodeBlogNotFound = 30001
)

// 定义错误
var (
	// General error
	InternalErr      = &BizErr{Code: CodeInternalErr, Msg: "未知错误"}
	InvalidParamErr  = &BizErr{Code: CodeInvalidParam, Msg: "无效参数"}
	UUIDGeneratorErr = &BizErr{Code: CodeUUIDGeneratorErr, Msg: "内部服务错误"}
	UserNotLogin     = &BizErr{Code: CodeNotLogin, Msg: "当前用户未登入"}
	UnmarshalErr     = &BizErr{Code: CodeUnmarshalErr, Msg: "反序列化错误"}
	// User module
	UserNotFound      = &BizErr{Code: CodeUserNotFoune, Msg: "用户不存在"}
	UserAlreadyExists = &BizErr{Code: CodeUserAleadyExists, Msg: "用户名已被注册"}

	// Blog module
	BlogNotFound = &BizErr{Code: CodeBlogNotFound, Msg: "博文不存在"}
)
