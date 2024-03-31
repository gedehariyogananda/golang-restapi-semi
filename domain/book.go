package domain

import (
	"test/golang/helper"
	"time"
)

type Book struct {
	Id          int64     `gorm:"primary_key"`
	Title       string    `gorm:"type:varchar(255);NOT NULL"`
	Author      string    `gorm:"type:varchar(255);NOT NULL"`
	Description string    `gorm:"type:varchar(255);NOT NULL"`
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type BookRepository interface {
	FindAll() ([]Book, error)
	AuthorExists(author string) error
	FindByAuthor(author string) ([]Book, error)
	Create(book Book) error
}

type BookService interface {
	FindAll() helper.ApiResponse
	FindByAuthor(author string) helper.ApiResponse
	Create(book Book) helper.ApiResponse
}

type BookData struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
