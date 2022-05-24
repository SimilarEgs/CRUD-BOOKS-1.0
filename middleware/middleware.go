package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/SimilarEgs/CURD-BOOKS/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//establishing a DB connection
func createDBConnection() *sql.DB {
	//loading environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error] .env file didn't load")
	}
	//open DB connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	//check connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[Info] DB was successfully connected")
	return db
}

func GetBookById() {}

func GetAllBooks() {}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	//getting book id from request params
	params := mux.Vars(r)
	//converting received id into int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("[Error] failed convertion string to int  %v", err)
	}
	//evoke «delete book» function
	deleteBook(int64(id))

	res := response{
		ID:      int64(id),
		Message: "[Info] book was successfully deleted from DB",
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateBookById() {}

//create book entity in DB
func CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Books
	//decode request body to variable book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatalf("[Error] cannot decode request body: %v", err)
	}

	//evoke «insert book» function
	insertID := insertBook(book)

	//preparing response object
	res := response{
		ID:      insertID,
		Message: "[Info] book was successfully created and stored in DB",
	}
	json.NewEncoder(w).Encode(res)
}

/////////////////////////////////////////////////////////////////////////
//                            handlers                          	  //
/////////////////////////////////////////////////////////////////////////

//insert book in the DB
func insertBook(book models.Books) int64 {

	db := createDBConnection()
	defer db.Close()

	//create INSERT sql query
	sqlStatement := `
	INSERT INTO books (bookname, author, date)
	VALUES ($1, $2, $3) 
	RETURNING id`

	//inserted id will store in this variable
	var id int64
	err := db.QueryRow(sqlStatement, book.BookName, book.Author, book.Date).Scan(&id)
	if err != nil {
		log.Fatalf("[Error] unable to execute the query: %v", err)
	}
	fmt.Printf("[Info] entity id - %v: was successfully inserted\n", id)

	return id
}

func deleteBook(id int64) {

	db := createDBConnection()
	defer db.Close()

	//create DELETE sql query
	sqlStatement := `
	DELETE FROM books
	WHERE id =$1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("[Error] failed to execute the query: %v", err)
	}

}
