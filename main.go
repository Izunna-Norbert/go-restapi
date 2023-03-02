package main

import (
	"encoding/json"
	"log"
	"net/http"

	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//BOOK MODEL
type Book struct {
	ID			string		`json:"id"`
	Isbn		string		`json:"isbn"`
	Title		string		`json:"title"`
	Author		*Author		`json:"author"`
}

type Author struct {
	Firstname		string		`json:"firstname"`
	Lastname		string		`json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			var book Book
			books = append(books[:index], books[index+1:]... )

			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func main () {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "First Book", Isbn: "1245362", Author: &Author{Firstname: "Authur", Lastname: "Pendragon"}})
	books = append(books, Book{ID: "2", Title: "Second Book", Isbn: "1245362", Author: &Author{Firstname: "Merlin", Lastname: "Magic"}})
	books = append(books, Book{ID: "3", Title: "Third Book", Isbn: "1245362", Author: &Author{Firstname: "Uther", Lastname: "Pendragon"}})
	// handle routes
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))

}