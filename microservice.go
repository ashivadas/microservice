package main

import (
	"fmt"
	"net/http"
	"cloudnative/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/",  api.BookHandleFunc)

	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0];
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Yo..Cloud Native Go.")
}