package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" example:"How To Become Backend Engineer"`
	Author string `json:"author" example:"John Smith"`
	Page   int    `json:"page" example:"205"`
}
