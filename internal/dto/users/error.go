package users

import "github.com/qhmd/gitforgits/internal/dto/book"

type ErrorResponse struct {
	*book.ErrorResponse
}

type InvalidId struct {
	*book.InvalidId
}

type UserNotFoundResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"user not found"`
	Errors  string `json:"errors" example:"user with id {id} does not exist"`
}

type EmailAlreadyUsed struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Email already exist, choose another Email"`
	Errors  string `json:"error" example:"this is email already used"`
}
