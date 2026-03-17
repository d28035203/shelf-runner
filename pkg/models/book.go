// Package models defines the Book entity and database access helpers (data layer).
package models

import (
	"errors"

	"github.com/d28035203/shelf-runner/pkg/config"
	"gorm.io/gorm"
)

// db is initialized once via init() so controllers can call helpers without passing a connection.
var db *gorm.DB

// Book is a catalog entry. gorm.Model embeds ID, CreatedAt, UpdatedAt, DeletedAt.
type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:255;not null"`
	Author      string `json:"author" gorm:"size:255;not null"`
	Publication string `json:"publication" gorm:"size:255"`
}

// init runs when the package is first imported: connect to MySQL and migrate the books table.
func init() {
	config.Connect()
	db = config.GetDB()
	_ = db.AutoMigrate(&Book{})
}

// CreateBook inserts the receiver as a new row and returns it (with ID populated).
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks returns every non-deleted book in the catalog.
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById loads a single book by primary key or returns "book not found".
func GetBookById(id int64) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, result.Error
	}
	return &book, nil
}

// UpdateBook persists all fields on an existing book (full save).
func UpdateBook(b *Book) error {
	return db.Save(b).Error
}

// DeleteBook soft-deletes a book by id (gorm.Model uses soft delete by default)
// and returns the record as it was before deletion when found.
func DeleteBook(id int64) Book {
	var book Book
	db.First(&book, id)
	db.Delete(&book)
	return book
}
