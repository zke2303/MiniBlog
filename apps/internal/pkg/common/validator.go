// Package common 常见的工具类
package common

import (
	"github.com/gin-gonic/gin/binding"
	chinese "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var GlobalTrans ut.Translator

func InitTranslation() {
	// 创建中文实例
	zh := chinese.New()
	// 创建翻译器, 并设置支持的语言
	uni := ut.New(zh, zh) // 设置默认语言为中文
	// 获取中文翻译
	trans, _ := uni.GetTranslator("zh")
	// 进行类型断言
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}
	// 注册中文翻译到 validator 中，由于 gin 中的数据校验使用了 validator，因此会使用这个翻译器
	if err := zhTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return
	}

	// 注册全局变量
	GlobalTrans = trans
}
