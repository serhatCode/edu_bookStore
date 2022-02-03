package main

import (
	"edu_bookStore/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
}
