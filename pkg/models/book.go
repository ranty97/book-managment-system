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

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var b []Book
	db.Find(&b)
	return b
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("ID=?", Id).Find(&b)

	return &b, db
}

func DeleteBook(Id int64) Book {
	var b Book
	db.Where("ID=?", Id).Delete(b)
	return b
}
