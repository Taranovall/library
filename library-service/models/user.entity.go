package models

import (
	"github.com/jinzhu/gorm"
	"library-service/utils"
	"time"
)

type UserEntity struct {
	gorm.Model
	Username string `gorm:"not null;unique_index:idx_username"`
	Password string
}

func (entity *UserEntity) BeforeCreate(db *gorm.DB) error {
	entity.Password = utils.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *UserEntity) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
