package initialize

import (
	"fmt"
	"go-api/global"
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
		zap.S().Errorf("init validator global.Trans failed %v\n", err)
		return
	}
	zap.S().Infof("Init Trans successfully")
}





// NOTE: locale can be retrieved from header "Accept-Language"
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// NOTE: use json key instead of entity key
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			zap.S().Infof("INitTrans %v", fld)
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
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}

	return
}

