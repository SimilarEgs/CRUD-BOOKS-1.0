package router

import (
	"github.com/SimilarEgs/CURD-BOOKS/middleware"
	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/book{id}", middleware.getBookById())
	router.HandleFunc("/book", middleware.getAllBooks())
	router.HandleFunc()
	router.HandleFunc()
	router.HandleFunc()

	return
}
