package utils

import "github.com/go-playground/validator/v10"

func MsgForTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return e.Field() + " must be a valid email address"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must be at most " + e.Param() + " characters"
	case "gt":
		return e.Field() + " must be greater than " + e.Param()
	case "lt":
		return e.Field() + " must be less than " + e.Param()
	case "alphaSpace":
		return e.Field() + " must only contain letters and spaces"
	case "alphaMin4":
		return e.Field() + " must be at least 4 characters"
	case "password":
		return e.Field() + " must be at least 8 characters, contain letters, numbers, symbols, and no spaces"
	default:
		return e.Field() + " is not valid"
	}
}
