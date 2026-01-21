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
		Code:  CodeInternal,
		Msg:   msg,
		Cause: cause,
	}
}

// 定义错误码
const (
	CodeInternal     = 10000
	CodeInvalidParam = 10001

	CodeUserNotFoune     = 20001
	CodePasswordError    = 20002
	CodeUserAleadyExists = 20003
)

// 定义错误
var (
	// General error
	InternalErr     = &BizErr{Code: CodeInternal, Msg: "Unknow error"}
	InvalidParamErr = &BizErr{Code: CodeInvalidParam, Msg: "Invalid params"}

	// User module
	UserNotFound      = &BizErr{Code: CodeUserNotFoune, Msg: "The User not found"}
	UserAlreadyExists = &BizErr{Code: CodeUserAleadyExists, Msg: "The user already exists."}
)
