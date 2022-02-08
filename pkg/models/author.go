package models

import (
	"edu_bookStore/pkg/config"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Books    []Book `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"books"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Author{})
}

func (author *Author) CreateAuthor() *Author {
	db.Create(&author)
	return author
}

func GetAllAuthors() []Author {
	var Authors []Author
	db.Find(&Authors)
	return Authors
}

func GetAuthorById(Id int64) (*Author, *gorm.DB) {
	var getAuthor Author
	db := db.Where("ID=?", Id).Find(&getAuthor)
	return &getAuthor, db
}
