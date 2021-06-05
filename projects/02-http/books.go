package main

import (
	"encoding/json"
	"io/ioutil"
)

type Book struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       string `json:"price"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

func getBooks() []Book {
	books := []Book{}
	contentBytes, err := ioutil.ReadFile("./books.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(contentBytes, &books)
	if err != nil {
		panic(err)
	}
	return books
}

func updateBooks(books []Book) {
	byteContent, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./books-updated.json", byteContent, 0777)
	if err != nil {
		panic(err)
	}
}
