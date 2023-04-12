package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Book struct {
	ID        int    `json:"id"`
	BookTitle string `json:"Book_title"`
	AuthorID  int    `json:"author_id"`
	ReaderID  int    `json:"reader_id"`
}

type Members struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

const (
	host     = "localhost"
	user     = "your_firstName_LastName"
	password = "your_password"
	dbName   = "your_dbName"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "mysql://user:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.HandleFunc("/authors", getAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", getAuthor).Methods("GET")
	r.HandleFunc("/authors", createAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", updateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", deleteAuthor).Methods("DELETE")

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	r.HandleFunc("/readers", getReaders).Methods("GET")
	r.HandleFunc("/readers/{id}", getReader).Methods("GET")
	r.HandleFunc("/readers", createReader).Methods("POST")
	r.HandleFunc("/readers/{id}", updateReader).Methods("PUT")
	r.HandleFunc("/readers/{id}", deleteReader).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", a))
}
func getAuthors(w http.ResponseWriter, r *http.Request) {
	authors := []Author{}

	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		author := Author{}
		err := rows.Scan(&author.Id, &author.FirstName, &author.LastName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	json.NewEncoder(w).Encode(authors)
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CreateAuthors(w http.ResponseWriter, r *http.Request) {
	authors := []Author{}

	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		author := Author{}
		err := rows.Scan(&author.Id, &author.FirstName, &author.LastName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	json.NewEncoder(w).Encode(authors)
}
