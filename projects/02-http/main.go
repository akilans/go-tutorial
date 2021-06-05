package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", HandleGetBooks)

	http.HandleFunc("/save", HandleSaveBooks)

	http.ListenAndServe(":8080", nil)
}

func HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	books := getBooks()
	// Marshal converts struct to json encoded [forming json]
	// Unmarshal converts json to Struct [parsing]
	byteContent, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	w.Write(byteContent)
}

func HandleSaveBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		fmt.Fprintf(w, r.Method+" is not allowed")
	} else {
		newBookByte, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		// to store new books
		var newBooks []Book

		// get old books
		books := getBooks()

		err = json.Unmarshal(newBookByte, &newBooks)

		if err != nil {
			panic(err)
		}

		// append the newbooks to the old books
		for _, newBook := range newBooks {
			books = append(books, newBook)
		}

		// Write all the books in new file
		updateBooks(books)

		w.Write([]byte("Success post"))
	}
}
