package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", HandleGetBooks)

	http.HandleFunc("/book", HandleGetBook)

	http.HandleFunc("/add", HandleAddBooks)

	http.HandleFunc("/delete", HandleDeleteBook)

	http.HandleFunc("/update", HandleUpdateBook)

	http.ListenAndServe(":8080", nil)
}

func HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	books := getBooks()
	// Marshal converts struct to json encoded [forming json]
	// Unmarshal converts json to Struct [parsing]
	byteContent, err := json.Marshal(books)
	panicError(err)
	w.Write(byteContent)
}

func HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")

	book, book_index := getBookById(id)

	if (Book{}) == book {
		w.Write(jsonMessageByte("Book Not found"))
	} else {
		books := getBooks()
		books = append(books[:book_index], books[book_index+1:]...)
		saveBooks(books)
		w.Write(jsonMessageByte("Book deleted successfully"))
	}

}

func HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " is not allowed"))
	} else {
		var updateBook Book

		updateBookByte, err := ioutil.ReadAll(r.Body)

		panicError(err)

		err = json.Unmarshal(updateBookByte, &updateBook)

		panicError(err)

		id := updateBook.Id

		book, _ := getBookById(id)

		if (Book{}) == book {
			w.Write(jsonMessageByte("Book Not found"))
		} else {
			books := getBooks()

			for i, book := range books {
				if book.Id == updateBook.Id {
					books[i] = updateBook
				}
			}
			saveBooks(books)
			w.Write(jsonMessageByte("Books saved successfully"))
		}

	}

}

func HandleAddBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " is not allowed"))
	} else {
		newBookByte, err := ioutil.ReadAll(r.Body)

		panicError(err)

		// to store new books
		var newBooks []Book

		// get old books
		books := getBooks()

		err = json.Unmarshal(newBookByte, &newBooks)

		panicError(err)

		// append the newbooks to the old books
		for _, newBook := range newBooks {
			books = append(books, newBook)
		}

		// Write all the books in new file
		saveBooks(books)

		w.Write(jsonMessageByte("Books saved successfully"))
	}
}

func HandleGetBook(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	id := query.Get("id")

	book, _ := getBookById(id)

	if (Book{}) == book {
		w.Write(jsonMessageByte("Book Not found"))
	} else {
		byteContent, err := json.Marshal(book)
		panicError(err)
		w.Write(byteContent)
	}

}
