package auth_controller

import (
	"library-service/models"
)

type Service interface {
	RegisterUser(input *UserInput) (*models.UserEntity, int)
	LogInUser(input *UserInput) (*models.UserEntity, int)
}

type service struct {
	repository Repository
}

func NewUserService(r Repository) *service {
	return &service{repository: r}
}

func (s service) RegisterUser(input *UserInput) (*models.UserEntity, int) {
	user := models.UserEntity{
		Username: input.Username,
		Password: input.Password}
	return s.repository.CreateUser(&user)
}

func (s service) LogInUser(input *UserInput) (*models.UserEntity, int) {

	user := models.UserEntity{
		Username: input.Username,
		Password: input.Password,
	}
	return s.repository.LogInUser(&user)
}
