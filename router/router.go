package router

import (
	"github.com/SimilarEgs/CURD-BOOKS/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/book/{id}", middleware.GetBookById).Methods("GET", "OPTIONS")
	router.HandleFunc("/book", middleware.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/book/{id}", middleware.DeleteBookById).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/book/{id}", middleware.UpdateBookById).Methods("PUT", "OPTIONS")
	router.HandleFunc("/book", middleware.CreateBook).Methods("POST", "OPTIONS")

	return router
}
