package book

import (
	"test/golang/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.BookRepository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.Book, error) {
	var books []domain.Book

	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) AuthorExists(author string) error {
	var book domain.Book

	if err := r.db.Where("author = ?", author).First(&book).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByAuthor(author string) ([]domain.Book, error) {
	var books []domain.Book

	if err := r.db.Where("author = ?", author).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) Create(book domain.Book) error {
	if err := r.db.Create(&book).Error; err != nil {
		return err
	}

	return nil
}
