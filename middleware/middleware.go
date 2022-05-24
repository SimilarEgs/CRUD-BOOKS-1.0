package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/SimilarEgs/CURD-BOOKS/db"
	"github.com/SimilarEgs/CURD-BOOKS/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//establishing a DB connection

func createDBConnection() *sql.DB{
	//loading environment file
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("[Error] .env file didn't load")
	}
	//open DB connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil{
		panic(err)
	}
	//check connection
	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("[Info] DB was successfully connected")
	return db
}



func getBookById(){}

func getAllBooks() {}

func deleteBookById() {}

func updateBookById() {}

func createBook() {}
