package user

import (
	"context"

	"github.com/qhmd/gitforgits/internal/domain/auth"
	authDto "github.com/qhmd/gitforgits/internal/dto/auth"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*auth.Auth, error)
	ListUser(ctx context.Context) ([]*auth.Auth, error)
	FindByEmail(ctx context.Context, email string) (*auth.Auth, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, users *authDto.UserResponse, id int) (*authDto.UserResponse, error)
}
