package database

import (
	"fmt"
	"time"

	"github.com/qhmd/gitforgits/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	dsn := config.GetDSN()
	var db *gorm.DB
	var err error
	// Retry sampai 5 kali
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to DB")
			return db
		}
		fmt.Println(dsn)
		fmt.Println("Waiting for DB to be ready...")
		time.Sleep(2 * time.Second)
	}
	panic("Failed to connect database, Error : " + err.Error())
}
