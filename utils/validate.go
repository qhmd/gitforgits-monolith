package utils

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	Validate.RegisterValidation("alphaSpace", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		regex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
		return regex.MatchString(value)
	})

	Validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		var hasLetter, hasNumber, hasSymbol, hasSpace bool
		for _, c := range password {
			switch {
			case c >= 'A' && c <= 'Z':
				hasLetter = true
			case c >= 'a' && c <= 'z':
				hasLetter = true
			case c >= '0' && c <= '9':
				hasNumber = true
			case c == ' ':
				hasSpace = true
			case (c >= 33 && c <= 47) || (c >= 58 && c <= 64) ||
				(c >= 91 && c <= 96) || (c >= 123 && c <= 126):
				hasSymbol = true
			}
		}

		return len(password) >= 8 && hasLetter && hasNumber && hasSymbol && !hasSpace
	})

	Validate.RegisterValidation("alphaMin4", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		trimmed := strings.ReplaceAll(value, " ", "")
		match := regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(value)
		return match && len(trimmed) >= 4
	})

}
