package database

import (
	"fmt"

	"github.com/qhmd/gitforgits/internal/domain/auth"
	"github.com/qhmd/gitforgits/internal/domain/book"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	if err := db.AutoMigrate(&book.Book{}, &auth.Auth{}); err != nil {
		fmt.Println("Migration gagal:", err)
		return
	}
	fmt.Println("Migration berhasil")
}
