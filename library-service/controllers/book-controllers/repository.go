package book_controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"library-service/models"
	"net/http"
)

type Repository interface {
	CreateBook(book *models.BookEntity) (*models.BookEntity, int, error)
	FindById(id uint) (*models.BookEntity, int, error)
	FindAllBooks() ([]models.BookEntity, int, error)
	DeleteBook(id uint) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) CreateBook(book *models.BookEntity) (*models.BookEntity, int, error) {
	createFile := repo.db.Create(&book)

	if err := createFile.Error; err != nil {
		logrus.Warnf(
			"An error occurred while creating the book. Status code: 500, Message: %s",
			err.Error(),
		)
		return nil, http.StatusInternalServerError, err
	}

	return book, http.StatusCreated, nil
}

func (repo *repository) FindById(id uint) (*models.BookEntity, int, error) {
	var book models.BookEntity

	if err := repo.db.First(&book, id).Error; err != nil {
		logrus.Warnf("Book with ID %d not found", id)
		return nil, http.StatusNotFound, err
	}

	return &book, http.StatusOK, nil
}

func (repo *repository) FindAllBooks() ([]models.BookEntity, int, error) {
	var books []models.BookEntity
	if err := repo.db.Find(&books).Error; err != nil {
		logrus.Warn(err.Error())
		return nil, http.StatusNotFound, err
	}
	return books, http.StatusOK, nil
}

func (repo *repository) DeleteBook(id uint) (int, error) {
	result := repo.db.Delete(&models.BookEntity{}, id)
	if result.RowsAffected == 0 {
		logrus.Infof("There's no book with ID: %d", id)
		return http.StatusNotFound, nil
	} else if err := result.Error; err != nil {
		logrus.Warn(err.Error())
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
