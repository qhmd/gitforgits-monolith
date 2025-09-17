package users

import "github.com/qhmd/gitforgits/internal/domain/auth"

type SuccessGetUser struct {
	Success bool       `json:"success" example:"true"`
	Message string     `json:"message" example:"successfully get the user"`
	Data    *auth.Auth `json:"data"`
}

type SuccessGetList struct {
	Success bool         `json:"success" example:"true"`
	Message string       `json:"message" example:"successfully get list user"`
	Data    []*auth.Auth `json:"data"`
}

type SuccessDeleteUser struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success delete user"`
}
