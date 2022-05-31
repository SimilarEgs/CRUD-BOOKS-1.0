package router

import (
	"github.com/SimilarEgs/CURD-BOOKS/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/book/{id}", middleware.GetBookById).Methods("GET"
	router.HandleFunc("/book", middleware.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", middleware.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/book/{id}", middleware.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book", middleware.CreateBook).Methods("POST")

	return router
}
