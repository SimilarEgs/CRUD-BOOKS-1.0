package router

import (
	"github.com/SimilarEgs/CURD-BOOKS/middleware"
	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/book/{id}", middleware.getBookById).Methods("GET")
	router.HandleFunc("/book", middleware.getAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", middleware.deleteBookById).Methods("DELETE")
	router.HandleFunc("book/{id}", middleware.updateBookById).Methods("PUT")
	router.HandleFunc("/book", middleware.updateBookById).Methods("POST")

	return
}
