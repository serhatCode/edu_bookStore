package controllers

import (
	"edu_bookStore/pkg/models"
	"encoding/json"
	"net/http"
)

func GetBookTypes(writer http.ResponseWriter, request *http.Request) {
	newBookType := models.GetAllBookTypes()
	res, _ := json.Marshal(newBookType)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(res)
	if err != nil {
		return
	}
}
