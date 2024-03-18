package book_controllers

import (
	"library-service/models"
)

type Service interface {
	CreateBook(book *BookInput) (*models.BookEntity, int, error)
	FindById(id uint) (*models.BookEntity, int, error)
	FindAllBooks() ([]models.BookEntity, int, error)
	DeleteBook(id uint) (int, error)
}

type service struct {
	repository Repository
}

func NewBookService(r Repository) *service {
	return &service{repository: r}
}

func (s service) CreateBook(book *BookInput) (*models.BookEntity, int, error) {
	bookEntity := models.BookEntity{
		Title:  book.Title,
		Author: book.Author,
	}
	return s.repository.CreateBook(&bookEntity)
}

func (s service) FindById(id uint) (*models.BookEntity, int, error) {
	return s.repository.FindById(id)
}

func (s service) FindAllBooks() ([]models.BookEntity, int, error) {
	return s.repository.FindAllBooks()
}

func (s service) DeleteBook(id uint) (int, error) {
	return s.repository.DeleteBook(id)
}
