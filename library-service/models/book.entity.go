package models

import "github.com/jinzhu/gorm"

type BookEntity struct {
	gorm.Model
	Title  string `gorm:"not null;unique_index:idx_title_author"`
	Author string `gorm:"not null;unique_index:idx_title_author"`
}
