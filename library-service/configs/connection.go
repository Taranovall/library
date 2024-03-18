package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"library-service/models"
	"os"
)

type DBConnection interface {
	Connection() *gorm.DB
}

type service struct{}

func NewDBService() *service {
	return &service{}
}

func (s *service) Connection() *gorm.DB {

	databaseURI := os.Getenv("DSN")

	db, err := gorm.Open("postgres", databaseURI)

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	} else {
		logrus.Info("Connection to Database Successfully")
	}

	databaseMigrations(db)

	return db
}

func databaseMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.BookEntity{}, &models.UserEntity{})
	logrus.Info("Database migrations")
}
