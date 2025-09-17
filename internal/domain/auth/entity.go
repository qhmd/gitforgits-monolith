package auth

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Name        string
	Email       string `gorm:"type:varchar(191)"`
	Role        string
	Password    string
	EmailUnique *string `gorm:"->;type:varchar(191) GENERATED ALWAYS AS (CASE WHEN deleted_at IS NULL THEN email ELSE NULL END) VIRTUAL;uniqueIndex" json:"-"`
}
