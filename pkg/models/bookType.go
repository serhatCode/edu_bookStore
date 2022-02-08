package models

import (
	"edu_bookStore/pkg/config"
	"gorm.io/gorm"
)

type BookType struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"foreignKey:bookTypeId, constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"books"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BookType{})
}

func (bookType *BookType) CreateBookType() *BookType {
	db.Create(&bookType)
	return bookType
}

func GetAllBookTypes() []BookType {
	var BookTypes []BookType
	db.Find(&BookTypes)
	return BookTypes
}
