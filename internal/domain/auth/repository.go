package auth

import (
	"context"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, auth *Auth) error
	UpdateMe(ctx context.Context, auth *Auth) (*Auth, error)
	GetUserByID(ctx context.Context, id uint) (*Auth, error)
	FindByEmail(ctx context.Context, email string) (*Auth, error)
	DeleteUser(ctx context.Context, id uint) error
	LogoutUser(ctx context.Context, token string) error
}
