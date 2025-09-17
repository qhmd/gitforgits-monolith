package usecase

import (
	"context"

	"github.com/qhmd/gitforgits/internal/domain/auth"
	"github.com/qhmd/gitforgits/internal/domain/user"
	authDto "github.com/qhmd/gitforgits/internal/dto/auth"
	"github.com/qhmd/gitforgits/utils"
)

type UsersUseCase struct {
	repo user.UserRepository
}

func NewUsersUseCase(repo user.UserRepository) *UsersUseCase {
	return &UsersUseCase{repo: repo}
}

func (uc *UsersUseCase) ListUser(ctx context.Context) ([]*auth.Auth, error) {
	return uc.repo.ListUser(ctx)
}

func (uc *UsersUseCase) GetUserByID(ctx context.Context, id int) (*auth.Auth, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UsersUseCase) UpdateUser(ctx context.Context, user *authDto.UserResponse, id int) (*authDto.UserResponse, error) {
	pw, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	data := &authDto.UserResponse{
		RegisterRequest: authDto.RegisterRequest{
			Name:     user.Name,
			Email:    user.Email,
			Password: pw,
		},
		Role: user.Role,
	}

	return uc.repo.UpdateUser(ctx, data, id)
}

func (uc *UsersUseCase) DeleteUser(ctx context.Context, id int) error {
	return uc.repo.DeleteUser(ctx, id)
}
