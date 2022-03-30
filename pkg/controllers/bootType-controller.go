package controllers

import (
	"edu_bookStore/pkg/models"
	"edu_bookStore/pkg/utils"
	"encoding/json"
	"net/http"
)

func GetBookTypes(writer http.ResponseWriter, request *http.Request) {
	newBookTypes := models.GetAllBookTypes()
	res, _ := json.Marshal(newBookTypes)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}

func CreateBookType(writer http.ResponseWriter, request *http.Request) {
	CreateBookType := &models.BookType{}
	utils.ParseBody(request, CreateBookType)
	bookType := CreateBookType.CreateBookType()
	res, _ := json.Marshal(bookType)
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}
