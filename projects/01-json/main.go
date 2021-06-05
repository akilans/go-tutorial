package main

import "fmt"

func main() {

	books := getBooks()

	for i, book := range books {
		fmt.Println(i+1, book.Title)
		books[i].Description = " Added for Marshal"
	}

	saveBooks(books)

}
