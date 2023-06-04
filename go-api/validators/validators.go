package validator

import (
	"github.com/go-playground/validator/v10"
)



func ValidateMobile(f validator.FieldLevel) bool {
  // mobile :=	f.Field().String()
	// NOTE: Taiwan cell phone
  // ok, _ := 	regexp.MatchString(`((?=(09))[0-9]{10})$`, mobile)
	return true
}