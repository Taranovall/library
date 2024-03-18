package auth_controller

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"library-service/models"
	"library-service/utils"
	"net/http"
	"strings"
)

type Repository interface {
	CreateUser(u *models.UserEntity) (*models.UserEntity, int)
	LogInUser(u *models.UserEntity) (*models.UserEntity, int)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(u *models.UserEntity) (*models.UserEntity, int) {
	checkUserAccount := r.db.Select("*").Where("username=?", u.Username).Find(&u)

	if checkUserAccount.RowsAffected > 0 {
		return nil, http.StatusConflict
	}

	createFile := r.db.Create(&u)

	if err := createFile.Error; err != nil {
		logrus.Warnf(
			"An error occurred while creating the user. Status code: 500, Message: %s",
			err.Error(),
		)
		return nil, http.StatusInternalServerError
	}

	return u, http.StatusCreated
}

func (r *repository) LogInUser(input *models.UserEntity) (*models.UserEntity, int) {
	var user models.UserEntity
	username := strings.ToLower(input.Username)

	if err := r.db.Where("LOWER(username) = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Warnf("UserEntity with username %s not found, error: %s", username, err)
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}

	comparePassword := utils.ComparePassword(user.Password, input.Password)

	if comparePassword != nil {
		return nil, http.StatusUnauthorized
	}

	return &user, http.StatusAccepted
}
