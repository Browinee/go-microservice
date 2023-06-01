package initialize

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"go.uber.org/zap"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	enTranslations "github.com/go-playground/validator/v10/translations/en"

	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)


func InitValidator() {
	if err := InitTrans("en"); err != nil {
		zap.S().Errorf("init validator trans failed %v\n", err)
		return
	}
}



var trans ut.Translator

// NOTE: locale can be retrieved from header "Accept-Language"
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// NOTE: use json key instead of entity key
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			fmt.Printf("INitTrans %v", fld)
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()

		// NOTE: first parameter is for fallback
		uni := ut.New(enT, zhT, enT)

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
