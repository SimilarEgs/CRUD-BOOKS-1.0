package main

import (
    "fmt"
    "log"
    "net/http"
  
  "github.com/SimilarEgs/CURD-BOOKS/router"
)

func main() {
  r := router.Router()
  log.Fatal(http.ListenAndServe(":8080", r))
  
    fmt.Println("Server started on port - 8080")
}
