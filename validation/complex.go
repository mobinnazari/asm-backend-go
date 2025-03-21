package validation

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func complexPassword(fl validator.FieldLevel) bool {
	var (
		special bool
		number  bool
		upper   bool
		lower   bool
	)

	for _, c := range fl.Field().String() {
		if special && number && upper && lower {
			break
		}
		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	if special && number && upper && lower {
		return true
	} else {
		return false
	}
}
