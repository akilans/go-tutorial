package main

import (
	"bytes"
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

func saveBooks(books []Book) {

	byteContent, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(byteContent))

	byteContent = bytes.Replace(byteContent, []byte("\\u003c"), []byte("<"), -1)
	byteContent = bytes.Replace(byteContent, []byte("\\u003e"), []byte(">"), -1)
	byteContent = bytes.Replace(byteContent, []byte("\\u0026"), []byte("&"), -1)

	//fmt.Println(string(byteContent))

	err = ioutil.WriteFile("./books-updated.json", byteContent, 0644)
	if err != nil {
		panic(err)
	}

}
