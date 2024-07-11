package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ranty97/book-managment-system-mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
