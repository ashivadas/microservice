package api

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Description string `json:"description, omitempty"`

}

var Books = map[string]Book {
	"sefn3494309": Book{Title: "The cloud", Author: "Bill B", ISBN: "sefn3494309"},
	"xxxxxx": Book{Title: "the sex", Author: "kamas", ISBN: "xxxxxx"}, 
}

// ToJSON to be used for marshalling
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if (err != nil) {
		panic(err)
	}
	return ToJSON
} 

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	fmt.Printf("%s", data)
	
	if (err != nil) {
		panic(err)
	}
	return book;
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
		case http.MethodGet:
			books := AllBooks();
			writeJSON(w, books);
		case http.MethodPost:
			body, err := ioutil.ReadAll(r.Body);
			if (err != nil) {
				w.WriteHeader(http.StatusInternalServerError)
			}
			book := FromJSON(body);
			isbn, created := CreateBook(book)
			if (created) {
				w.Header().Add("Location", "/api/books"+isbn)
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusConflict);
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unsupported request method"))
	}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
		case http.MethodGet:
			found, book := GetBook(isbn)
			if (found) {
				a := make ([]Book, 0)
				a = append (a, book)
				writeJSON(w, a)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		case http.MethodPut:
			body, err := ioutil.ReadAll(r.Body)
			if (err != nil) {
				w.WriteHeader(http.StatusInternalServerError)
			}
			book := FromJSON(body);
			isbn, updated := UpdateBook(book)
			if (updated) {
				w.Header().Add("Updated", "/api/books"+isbn)
				w.WriteHeader(http.StatusFound)
			} else {
				w.WriteHeader(http.StatusConflict)
			}
			
		case http.MethodDelete:
			ok := DeleteBook(isbn)
			if (ok) {
				w.WriteHeader(http.StatusOK)				
			}	else {
				w.WriteHeader(http.StatusNotFound)				
			}	
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unsupported request method"))
	}
}

func writeJSON(w http.ResponseWriter, b []Book) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	books, err := json.Marshal(b)
	if (err != nil) {
		panic(err)
	}
	w.Write(books)
}

func AllBooks() []Book{
	books := []Book{}
	for _,value := range Books {
		books = append(books, value)
	}
	return books;
}

func CreateBook(book Book) (string, bool) {
	_,ok := Books[book.ISBN]
	if !ok {
		Books[book.ISBN] = book
	}
	return book.ISBN, ok
}

func UpdateBook(book Book) (string, bool) {
	lookup := Book{}
	_,ok := Books[book.ISBN]
	if (ok) {
		lookup.Author = book.Author
		lookup.Title = book.Title
		lookup.Description = book.Description
		Books[book.ISBN] = lookup
	}
	return book.ISBN, ok
}

func GetBook(isbn string) (bool, Book) {
	book,ok := Books[isbn]
	return ok, book;
}

func DeleteBook(isbn string) (bool) {
	_,ok := Books[isbn]
	if (ok) {
		delete (Books, isbn)
	}
	return ok
}