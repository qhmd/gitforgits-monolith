package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/qhmd/gitforgits/config"
	"github.com/qhmd/gitforgits/internal/domain/auth"
	authDto "github.com/qhmd/gitforgits/internal/dto/auth"
	"github.com/qhmd/gitforgits/utils"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	repo auth.AuthRepository
}

func NewAuthUsecase(repo auth.AuthRepository) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (u *AuthUseCase) RegisterUser(ctx context.Context, us *auth.Auth) error {
	existing, err := u.repo.FindByEmail(ctx, us.Email)
	fmt.Printf("isi exis : %v, dan isi err %v", existing, err)
	if err != nil {
		return err
	}

	if existing != nil {
		return config.ErrUserExists
	}
	return u.repo.RegisterUser(ctx, us)
}

func (u *AuthUseCase) LoginUser(ctx context.Context, email, password string) (*auth.Auth, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (u *AuthUseCase) Me(ctx context.Context, email string) (*authDto.RegisterRequest, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	userReponse := &authDto.RegisterRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return userReponse, nil
}

func (u *AuthUseCase) UpdateMe(ctx context.Context, user *auth.Auth) (*authDto.RegisterRequest, error) {
	req, err := u.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, errors.New("user not found")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userUpdateReponse := &auth.Auth{
		Model:    gorm.Model{ID: user.ID},
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	putData, err := u.repo.UpdateMe(ctx, userUpdateReponse)
	if err != nil {
		return nil, err
	}

	dataToRegist := &authDto.RegisterRequest{
		Name:     putData.Name,
		Email:    putData.Email,
		Password: putData.Password,
	}

	return dataToRegist, nil
}

func (u *AuthUseCase) DeleteUserByID(ctx context.Context, id uint) (*auth.Auth, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *AuthUseCase) GetUserByID(ctx context.Context, id uint) (*auth.Auth, error) {
	return u.repo.GetUserByID(ctx, id)
}
