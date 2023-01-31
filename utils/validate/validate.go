package validate

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"goblog/utils/errmsg"
	"reflect"
)

// Validate 数据验证和翻译
func Validate(data any) (string, int) {
	validate := validator.New()
	uni := ut.New(zh_Hans_CN.New())

	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	err = validate.Struct(data) //断言并进行数据验证

	if err != nil { //数据验证不通过
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
