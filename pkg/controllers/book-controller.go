package controllers

import (
	"edu_bookStore/pkg/models"
	"edu_bookStore/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		return
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(request, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		return
	}
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	if updateBook.AuthorId != 0 {
		bookDetails.AuthorId = updateBook.AuthorId
	}

	if updateBook.BookTypeId != 0 {
		bookDetails.BookTypeId = updateBook.BookTypeId
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		return
	}
}
