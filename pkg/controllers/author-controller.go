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

func GetAuthors(writer http.ResponseWriter, request *http.Request) {
	newAuthor := models.GetAllAuthors()
	res, _ := json.Marshal(newAuthor)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}

func GetAuthorById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	authorId := vars["authorId"]
	Id, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	authorDetails, _ := models.GetAuthorById(Id)
	res, _ := json.Marshal(authorDetails)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		return
	}
}

func CreateAuthor(writer http.ResponseWriter, request *http.Request) {
	CreateAuthor := &models.Author{}
	utils.ParseBody(request, CreateAuthor)
	author := CreateAuthor.CreateAuthor()
	res, _ := json.Marshal(author)
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}

func UpdateAuthor(writer http.ResponseWriter, request *http.Request) {
	var updateAuthor = &models.Author{}
	utils.ParseBody(request, updateAuthor)
	vars := mux.Vars(request)
	authorId := vars["authorId"]
	Id, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	authorDetails, db := models.GetAuthorById(Id)
	if updateAuthor.Name != "" {
		authorDetails.Name = updateAuthor.Name
	}

	if updateAuthor.LastName != "" {
		authorDetails.LastName = updateAuthor.LastName
	}

	db.Save(&authorDetails)
	res, _ := json.Marshal(authorDetails)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		return
	}
}
