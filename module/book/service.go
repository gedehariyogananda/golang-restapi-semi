package book

import (
	"net/http"
	"test/golang/domain"
	"test/golang/helper"
	"time"
)

type service struct {
	bookRepository domain.BookRepository
}

func NewService(bookRepository domain.BookRepository) domain.BookService {
	return &service{bookRepository}
}

// FindAll implements domain.BookService.
func (s *service) FindAll() helper.ApiResponse {
	books, err := s.bookRepository.FindAll()
	if err != nil {
		return helper.ApiResponse{
			Message: "Failed to get books",
			Status:  false,
		}
	}

	return helper.ApiResponse{
		Message: "Success",
		Data:    books,
		Status:  true,
	}
}

// FindByAuthor implements domain.BookService.
func (s *service) FindByAuthor(author string) helper.ApiResponse {
	err := s.bookRepository.AuthorExists(author)
	if err != nil {
		return helper.ApiResponse{
			Message: "author not found",
			Status:  false,
		}
	}

	book, err := s.bookRepository.FindByAuthor(author)
	if err != nil {
		return helper.ApiResponse{
			Message: "Failed to get book",
			Status:  false,
			Code:    http.StatusInternalServerError,
		}
	}

	return helper.ApiResponse{
		Message: "Success",
		Data:    book,
		Status:  true,
		Code:    http.StatusOK,
	}

}

// Create implements domain.BookService.
func (s *service) Create(book domain.Book) helper.ApiResponse {
	if err := s.bookRepository.Create(book); err != nil {
		return helper.ApiResponse{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusInternalServerError,
		}
	}

	result := domain.BookData{
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		CreatedAt:   book.CreatedAt.Format(time.RFC3339),
	}

	return helper.ApiResponse{
		Message: "berhasil di add",
		Data:    result,
		Status:  true,
		Code:    http.StatusCreated,
	}

}
