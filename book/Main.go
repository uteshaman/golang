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
	db, err = sql.Open("postgres", "postgresql://user:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	
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

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
	authors := []Author{}
	rows, err := db.Query("SELECT id, name, pseudonym, specialty FROM authors")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()
	for rows.Next() {
		author := Author{}
		err := rows.Scan(&author.ID, &author.Name, &author.Pseudonym, &author.Specialty)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
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

	row:= db.QueryRow("SELECT * FROM authors WHERE id = $1", id)

	author := Author{}
	err = row.Scan(&author.Id, &author.FirstName, &author.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(author)

	func createAuthor(w http.ResponseWriter, r *http.Request){
     params := mux.Vars(r)
	 id, err := strconv.Atoi(params["id"])
	 if err != nil{
		http.Error(w err.Error(), http.StatusBadRequest)
		return
	 }
	}
	row := db.QueryRow("SELECT * FROM authors WHERE id = $1", id)
	author := Author{}
	err = row.Scan(&author.Id, &author.FirstName, &author.LastName)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(author)

	func updateAuthor(w http.ResponseWriter, r*http.Request){
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil{
		   http.Error(w err.Error(), http.StatusBadRequest)
		   return
		}
	   }
	   row := db.QueryRow("SELECT * FROM authors WHERE id = $1", id)
	   author := Author{}
	   err = row.Scan(&author.Id, &author.FirstName, &author.LastName)
	   if err != nil{
		   http.Error(w, err.Error(), http.StatusInternalServerError)
		   return
	   }
	   json.NewEncoder(w).Encode(author)


	func deleteAuthor(w http.ResponseWriter, r*http.Request){
		params := mux.Vars(r)
	 id, err := strconv.Atoi(params["id"])
	 if err != nil{
		http.Error(w err.Error(), http.StatusBadRequest)
		return
	 }
	}
	row := db.QueryRow("SELECT * FROM authors WHERE id = $1", id)
	author := Author{}
	err = row.Scan(&author.Id, &author.FirstName, &author.LastName)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(author)


	func getBooks(w http.ResponseWriter, r*http.Request){
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil{
		   http.Error(w err.Error(), http.StatusBadRequest)
		   return
		}
	   }
	   row := db.QueryRow("SELECT * FROM authors WHERE id = $1", id)
	   author := Author{}
	   err = row.Scan(&book.ID, &book.BookTitle,  &book.ReaderID, &book.AuthorID)
	   if err != nil{
		   http.Error(w, err.Error(), http.StatusInternalServerError)
		   return
	   }
	   json.NewEncoder(w).Encode(author)


	func getBook(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := db.QueryRow("SELECT * FROM books WHERE id = $1", id)

		book := Book{}
		err = row.Scan(&book.ID, &book.BookTitle,  &book.ReaderID, &book.AuthorID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(book)
	}

	func createBook(w http.ResponseWriter, r *http.Request) {
		book := Book{}

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO books (title, genre, isbn, author_id) VALUES ($1, $2, $3, $4)", &book.BookTitle,  &book.ReaderID, &book.AuthorID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		book.ID = int(id)

		json.NewEncoder(w).Encode(book)
	}
	
	func updateBook(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		book := Book{}
	
		err = json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		_, err = db.Exec("UPDATE books SET title = $1, genre = $2, isbn = $3, author_id = $4 WHERE id = $5",&book.BookTitle,  &book.ReaderID, &book.AuthorID, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		json.NewEncoder(w).Encode(book)
	}
	
	func deleteBook(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		_, err = db.Exec("DELETE FROM books WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		w.WriteHeader(http.StatusOK)
	}
	func getReaders(w http.ResponseWriter, r *http.Request) {
		readers := []Reader{}
	
		rows, err := db.Query("SELECT * FROM readers")
		if err != nil {	http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
	
		for rows.Next() {
			reader := Reader{}
			err := rows.Scan(&reader.ID, &reader.FullName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			
			readerBooksRows, err := db.Query("SELECT book_id FROM reader_books WHERE reader_id = $1", reader.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer readerBooksRows.Close()
	
			for readerBooksRows.Next() {
				var bookID int
				err := readerBooksRows.Scan(&bookID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
	
				bookRow := db.QueryRow("SELECT title, genre, isbn, author_id FROM books WHERE id = $1", bookID)
	
				book := Book{}
				err = bookRow.Scan(&book.Title, &book.Genre, &book.ISBN, &book.AuthorID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
	
				reader.Books = append(reader.Books, book)
			}
	
			readers = append(readers, reader)
		}
	
		json.NewEncoder(w).Encode(readers)
	}
	
	func getReader(w.http.ResponseWriter, r*http.Request){
reader := Reader{0

err := json.NewDecoder(r.body).Decode(&reader)
if err != nil{
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
			}
		}
	}
	
	func createReader(w http.ResponseWriter, r *http.Request) {
		reader := Reader{}
	
		err := json.NewDecoder(r.Body).Decode(&reader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		result, err := tx.Exec("INSERT INTO readers (full_name) VALUES ($1)", reader.FullName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		readerID, err := getReaderID(result)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		for _, book := range reader.Books {
			_, err := tx.Exec("INSERT INTO reader_books (reader_id, book_id) VALUES ($1, $2)", readerID, book.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		
		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		reader.ID = readerID
		json.NewEncoder(w).Encode(reader)
	}

 func updateReader(w http.ResponseWriter, r*http.Request){
 params := mux.Vars(r)
 id, err := strconv.Atoi(params["id"])
 if err != nil{
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
 	}
 }
 
	func deleteReader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
		}
	}
	