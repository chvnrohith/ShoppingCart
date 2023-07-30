package routes

import (
	"example.com/hello/Desktop/go-workspace/pkg/mod/github.com/!akhil!sharma90/!golang-!my!s!q!l-!c!r!u!d-!bookstore-!management-!a!p!i@v0.0.0-20211006083836-71e755118f4f/pkg/controllers"
	"example.com/hello/Desktop/go-workspace/pkg/mod/github.com/gorilla/mux@v1.8.0"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
