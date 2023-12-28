package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{"This is index page. If you see that, you can run my API server. Congratulation!"})
}

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Понедельник начинается в субботу", Author: &Author{
		FirstName: "Аркадий и Борис",
		LastName:  "Стругацкие",
	}})

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
