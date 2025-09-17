package repository

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/qhmd/gitforgits/config"
	"github.com/qhmd/gitforgits/internal/domain/auth"
	"github.com/qhmd/gitforgits/internal/domain/user"
	authDto "github.com/qhmd/gitforgits/internal/dto/auth"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// DeleteUser implements user.UserRepository.
func (m *mysqlUserRepository) DeleteUser(ctx context.Context, id int) error {
	return m.db.WithContext(ctx).Delete(&auth.Auth{}, id).Error
}

// FindByEmail implements user.UserRepository.
func (m *mysqlUserRepository) FindByEmail(ctx context.Context, email string) (*auth.Auth, error) {
	var user *auth.Auth
	if err := m.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser implements user.UserRepository.
func (m *mysqlUserRepository) GetUser(ctx context.Context, id int) (*auth.Auth, error) {
	var user *auth.Auth
	if err := m.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ListUser implements user.UserRepository.
func (m *mysqlUserRepository) ListUser(ctx context.Context) ([]*auth.Auth, error) {
	var users []*auth.Auth
	if err := m.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser implements user.UserRepository.
func (m *mysqlUserRepository) UpdateUser(ctx context.Context, users *authDto.UserResponse, id int) (*authDto.UserResponse, error) {
	if err := m.db.WithContext(ctx).Model(&auth.Auth{}).Where("id = ?", id).Updates(&users).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil, config.ErrUserExists
			}
		}
		return nil, err
	}
	fmt.Print("isi dari user di mysql", users)
	return users, nil
}

func NewUserMySqlRepo(db *gorm.DB) user.UserRepository {
	return &mysqlUserRepository{db: db}
}
