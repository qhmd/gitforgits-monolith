package auth

import "github.com/qhmd/gitforgits/internal/dto/book"

type ErrorResponseAuth struct {
	book.ErrorResponse
}

type ErrorResponseLogin struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"try another email"`
	Error   string `json:"error" example:"email already used"`
}

type ErrorUnauthorized struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"missing refresh token"`
}
