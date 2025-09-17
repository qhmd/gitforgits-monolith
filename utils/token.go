package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte(os.Getenv("ACCESS_SECRET"))
var refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

func GenerateAccessToken(id uint, email, name string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

func GenerateRefreshToken(id uint, email, name string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(14 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

func VerifyAccessToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (any, error) {
		return accessSecret, nil
	})
}

func VerifyRefreshToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (any, error) {
		return refreshSecret, nil
	})
}
