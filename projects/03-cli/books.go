package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

type Message struct {
	Msg string
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}

}

func jsonMessageByte(msg string) []byte {
	errrMessage := Message{msg}
	byteContent, err := json.Marshal(errrMessage)
	panicError(err)
	return byteContent
}

func getBooks() []Book {
	books := []Book{}
	contentBytes, err := ioutil.ReadFile("./books.json")

	panicError(err)

	err = json.Unmarshal(contentBytes, &books)
	panicError(err)
	return books
}

func saveBooks(books []Book) {
	byteContent, err := json.Marshal(books)

	panicError(err)

	err = ioutil.WriteFile("./books.json", byteContent, 0777)
	fmt.Println("Saved all books")
	panicError(err)
}

func getBookById(id string) (Book, int) {

	books := getBooks()
	var requestedBook Book
	var requestedBookIndex int

	for i, book := range books {
		if book.Id == id {
			requestedBook = book
			requestedBookIndex = i
		}
	}

	return requestedBook, requestedBookIndex

}
