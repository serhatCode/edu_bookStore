package routes

import (
	"edu_bookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/author/", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/", controllers.GetAuthors).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.GetAuthors).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.GetAuthors).Methods("PUT")
}
