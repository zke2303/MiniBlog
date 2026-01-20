// Package common 常见的工具类
package common

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	chinese "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var GlobalTrans ut.Translator

func InitTranslation() error {
	// 创建中文实例
	zh := chinese.New()
	// 创建翻译器, 并设置支持的语言
	uni := ut.New(zh, zh) // 设置默认语言为中文
	// 获取中文翻译
	trans, _ := uni.GetTranslator("zh")
	// 进行类型断言
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil
	}
	// 注册一个函数，获取 struct tag 里自定义的 label 作为字段名
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 注册中文翻译到 validator 中，由于 gin 中的数据校验使用了 validator，因此会使用这个翻译器
	if err := zhTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return err
	}

	// 注册全局变量
	GlobalTrans = trans
	return nil
}
