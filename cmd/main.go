package main

import (
	"books-api/models"
	"books-api/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	services.Books = append(services.Books, models.Book{ID: "1", Title: "Понедельник начинается в субботу", Author: &models.Author{
		FirstName: "Аркадий и Борис",
		LastName:  "Стругацкие",
	}})
	services.Books = append(services.Books, models.Book{ID: "2", Title: "Пикник на обочине", Author: &models.Author{
		FirstName: "Аркадий и Борис",
		LastName:  "Стругацкие",
	}})

	r.HandleFunc("/", services.Index).Methods("GET")
	r.HandleFunc("/books", services.GetBooks).Methods("GET")
	r.HandleFunc("/books", services.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", services.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", services.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", services.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
