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

	
	r.HandleFunc("/members", getMembers).Methods("GET")
	r.HandleFunc("/members/{id}", getMember).Methods("GET")
	r.HandleFunc("/members", createMember).Methods("POST")
	r.HandleFunc("/members/{id}", updateMember).Methods("PUT")
	r.HandleFunc("/members/{id}", deleteMember).Methods("DELETE")

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
		err := rows.Scan(&author.Id, &author.FirstName, &author.LastName)
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
	func getMembers(w http.ResponseWriter, r *http.Request) {
		members := []Members{}
	
		rows, err := db.Query("SELECT * FROM members")
		if err != nil {	http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
	
		for rows.Next() {
			reader := Members{}
			err := rows.Scan(&member.ID, &member.FullName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			
			membersBooksRows, err := db.Query("SELECT book_id FROM members_books WHERE reader_id = $1", reader.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer membersBooksRows.Close()
	
			for membersBooksRows.Next() {
				var bookID int
				err := membersBooksRows.Scan(&bookID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
	
				bookRow := db.QueryRow("SELECT title, genre, isbn, author_id FROM books WHERE id = $1", bookID)
	
				book := Book{}
				err = bookRow.Scan(&members.ID, &members.FirstName, &members.LastName)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
	
				member.Books = append(Members.Books, book)
			}
	
			members = append(members, member)
		}
	
		json.NewEncoder(w).Encode(members)
	}
	
	func getMember(w http.ResponseWriter, r*http.Request){
       member := Member{}

err := json.NewDecoder(r.body).Decode(&member)
if err != nil{
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
			}
		}
	}
	
	func createMember(w http.ResponseWriter, r *http.Request) {
		member := Members{}
	
		err := json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		result, err := tx.Exec("INSERT INTO members (full_name) VALUES ($1)", member.FullName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		memberID, err := getMemberID(result)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		for _, book := range member.Books {
			_, err := tx.Exec("INSERT INTO member_books (member_id, book_id) VALUES ($1, $2)", member.ID, book.ID)
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
		
		member.ID = memberID
		json.NewEncoder(w).Encode(member)
	}

 func updateMember(w http.ResponseWriter, r*http.Request){
 params := mux.Vars(r)
 id, err := strconv.Atoi(params["id"])
 if err != nil{
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
 	}
 }
 
	func deleteMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
		}
	}
	