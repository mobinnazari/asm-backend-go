package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var publicServices []string = []string{"gmail.com", "yahoo.com", "outlook.com"}

func nonPublicEmail(fl validator.FieldLevel) bool {
	res := true
	for _, service := range publicServices {
		if strings.Contains(fl.Field().String(), service) {
			res = false
			break
		}
	}
	return res
}
