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

//establishing DB connection
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
	return db
}

//return single book entity by its id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	//getting book id from request params
	params := mux.Vars(r)
	//converting received id into int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("[Error] failed convertion string to int  %v", err)
	}

	//check if row exists, if true -> call getbook function and send json
	if rowExists("SELECT id FROM books WHERE id=$1", id) {

		book, err := getbook(int64(id))
		if err != nil {
			log.Fatalf("[Error] unable to get book entity: %v", err)
		}
		json.NewEncoder(w).Encode(book)

	} else {
		book, _ := getbook(int64(id))
		json.NewEncoder(w).Encode(book)

	}
}

//return all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := getAllBooks()

	if err != nil {
		log.Fatalf("[Error] failed to get all books: %v", err)
	}

	json.NewEncoder(w).Encode(books)
}

//delete book entity from db
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	//getting book id from request params
	params := mux.Vars(r)

	//converting received id into int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("[Error] failed convertion string to int  %v", err)
	}

	//check if row exists, if true -> delete this book
	if rowExists("SELECT id FROM books WHERE id=$1", id) {

		deleteBook(int64(id)) //call delete function
		res := response{
			ID:      int64(id),
			Message: fmt.Sprintf("[Info] entity id - %v: was successfully deleted", id),
		}
		json.NewEncoder(w).Encode(res)

	} else {

		res := response{
			ID:      int64(id),
			Message: fmt.Sprintf("[Error] entity with id - %d: doesnt exists", id),
		}
		json.NewEncoder(w).Encode(res)
	}

}

//update book entity detail in db
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)

	//converting received id into int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("[Error] failed convertion string to int  %v", err)
	}
	
	var book models.Books
	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("[Error] cannot decode request body: %v", err)
	}
	//check if row exists, if true -> call update function to update book entity
	if rowExists("SELECT id FROM books WHERE id=$1", id) {

		updateBookById(int64(id), book)

		//format response object
		res := response{
			ID:      int64(id),
			Message: fmt.Sprintf("[Info] entity with id - %d: was successfully updated", id),
		}

		json.NewEncoder(w).Encode(res)

	} else {

		res := response{
			ID:      int64(id),
			Message: fmt.Sprintf("[Error] entity with id - %d: doesnt exists", id),
		}

		json.NewEncoder(w).Encode(res)
	}
}

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

	//format response object
	res := response{
		ID:      insertID,
		Message: "[Info] book was successfully created and stored in DB",
	}
	json.NewEncoder(w).Encode(res)
}

/////////////////////////////////////////////////////////////////////////
//                            handlers                                 //
/////////////////////////////////////////////////////////////////////////

func insertBook(book models.Books) int64 {

	db := createDBConnection()
	defer db.Close()

	var id int64

	//create INSERT sql statement
	sqlStatement := `
	INSERT INTO books (bookname, author, date)
	VALUES ($1, $2, $3) 
	RETURNING id`

	// execute the sql statement
	err := db.QueryRow(sqlStatement, book.BookName, book.Author, book.Date).Scan(&id)

	if err != nil {
		log.Fatalf("[Error] unable to execute sqlStatement: %v", err)
	}
	fmt.Printf("[Info] entity with id - %v: was successfully inserted\n", id)

	return id 
}

func getbook(id int64) (models.Books, error) {

	db := createDBConnection()
	defer db.Close()

	var book models.Books

	//create SELECT sql statement
	sqlStatement := `
	SELECT * FROM books
	WHERE id=$1`

	//execute SELECT sql statement
	row := db.QueryRow(sqlStatement, id)

	//unmarshaling row
	err := row.Scan(&book.ID, &book.BookName, &book.Author, &book.Date)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("[Error] now row was returned - ", sql.ErrNoRows)
		return book, nil
	case nil:
		return book, nil
	default:
		log.Fatalf("[Error] failed to scan the row %v", err)
	}
	//return empty book on error
	return book, err
}

func getAllBooks() ([]models.Books, error) {

	db := createDBConnection()
	defer db.Close()

	var books []models.Books

	//create SELECT sql statement
	sqlStatement := `SELECT * FROM books`

	//execute sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("[Error] unable to execute sqlStatement: %v", err)
	}
	//close sqlStatement
	defer rows.Close()

	//iterate over rows, append books variable
	for rows.Next() {
		var book models.Books

		//unmurshaling row
		err = rows.Scan(&book.ID, &book.BookName, &book.Author, &book.Date)

		if err != nil {
			log.Fatalf("[Error] failed to scan the row %v", err)
		}

		books = append(books, book)
	}

	//return empty book on error
	return books, err
}

func deleteBook(id int64) {

	db := createDBConnection()
	defer db.Close()

	//create DELETE sql statement
	sqlStatement := `
	DELETE FROM books
	WHERE id =$1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("[Error] failed to execute the query: %v", err)
	}

}

func updateBookById(id int64, book models.Books) {

	db := createDBConnection()
	defer db.Close()

	//create UPDATE sql statement
	sqlstatement := `
	UPDATE books 
	SET bookname = $2, author = $3, date = $4
	WHERE id = $1;`

	//execute the sql statement
	_, err := db.Exec(sqlstatement, id, book.BookName, book.Author, book.Date)
	if err != nil {
		log.Fatalf("[Error] failed to execute the query: %v", err)
	}
}

/////////////////////////////////////////////////////////////////////////
//                            healpers                                 //
/////////////////////////////////////////////////////////////////////////

func rowExists(query string, args ...interface{}) bool {

	db := createDBConnection()
	defer db.Close()

	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("[Error] '%s'%v", args, err)
	}
	return exists
}
