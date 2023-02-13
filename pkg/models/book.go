package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uint   `gorm:"primay key;autoincrement" json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func InitBook(newDb *gorm.DB) {
	newDb.AutoMigrate(&Book{})
	db = newDb
}

func (b *Book) CreateBook() (*Book, error) {
	err := db.Create(&b).Error
	return b, err
}

func GetAllBooks() ([]Book, error) {
	var bookModels []Book
	err := db.Find(&bookModels).Error
	return bookModels, err
}

func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var getBook Book
	err := db.Where("id = ?", Id).First(&getBook).Error
	return &getBook, db, err

}

func GetBookByNameAndAuthor(name string, author string) ([]Book, *gorm.DB, error) {

	query := `SELECT * FROM books where books.name=? and books.author=?`
	var results []Book
	err := db.Raw(query, name, author).Scan(&results).Error

	return results, db, err

}

func DeleteBook(Id int64) (Book, error) {
	var book Book
	err := db.Where("ID=?", Id).Delete(&book).Error

	return book, err
}
