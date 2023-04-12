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

	a.POST("/authors", createAuthor)
	a.GET("/authors", getAuthors)
	a.GET("/authors/:id", getAuthor)
	a.PUT("/authors/:id", updateAuthor)
	a.DELETE("/authors/:id", deleteAuthor)

	a.POST("/books", createBook)
	a.GET("/books", getBooks)
	a.GET("/books/:id", getBook)
	a.PUT("/books/:id", updateBook)
	a.DELETE("/books/:id", deleteBook)

	a.POST("/members", createMember)
	a.GET("/members", getMembers)
	a.GET("/members/:id", getMember)
	a.PUT("/members/:id", updateMember)
	a.DELETE("/members/:id", deleteMember)

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
		err := rows.Scan(&author.ID, &author.Name, &author.Pseudonym, &author.Specialization)
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
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
