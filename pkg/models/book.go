package models

import (
	"example.com/hello/Desktop/go-workspace/pkg/mod/github.com/!akhil!sharma90/!golang-!my!s!q!l-!c!r!u!d-!bookstore-!management-!a!p!i@v0.0.0-20211006083836-71e755118f4f/pkg/config"
	"example.com/hello/Desktop/go-workspace/pkg/mod/github.com/jinzhu/gorm@v1.9.12"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
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
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
