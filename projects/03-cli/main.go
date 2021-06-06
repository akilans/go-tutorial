package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all books")
	getId := getCmd.String("id", "", "Book id")

	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delId := delCmd.String("id", "", "Book id")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addId := addCmd.String("id", "", "Book id")
	addTitle := addCmd.String("title", "", "Book title")
	addAuthor := addCmd.String("author", "", "Book author")
	addPrice := addCmd.String("price", "", "Book price")
	addImageUrl := addCmd.String("imageurl", "", "Book image URL")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.String("id", "", "Book id")
	updateTitle := updateCmd.String("title", "", "Book title")
	updateAuthor := updateCmd.String("author", "", "Book author")
	updatePrice := updateCmd.String("price", "", "Book price")
	updateImageUrl := updateCmd.String("imageurl", "", "Book image URL")

	if len(os.Args) < 2 {
		fmt.Println("Expected get | add | update | delete argument needed")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGet(getCmd, getAll, getId)
	case "delete":
		handleDelete(delCmd, delId)
	case "update":
		handleUpdate(updateCmd, updateId, updateTitle, updateAuthor, updatePrice, updateImageUrl)
	case "add":
		handleAdd(addCmd, addId, addTitle, addAuthor, addPrice, addImageUrl)
	default:
		fmt.Println("Not sure what do to")
	}
}

func handleDelete(delCmd *flag.FlagSet, id *string) {
	delCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("Pass id option")
		delCmd.PrintDefaults()
		os.Exit(1)
	} else {
		book, book_index := getBookById(*id)
		if (Book{}) == book {
			fmt.Println("Book not found")
		} else {
			books := getBooks()
			books = append(books[:book_index], books[book_index+1:]...)
			saveBooks(books)
			fmt.Println("Book deleted successfully")
		}
	}
	return

}

func handleUpdate(updateCmd *flag.FlagSet, id *string, title *string, author *string, price *string, imageurl *string) {

	updateCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *author == "" || *price == "" || *imageurl == "" {
		fmt.Println("Pass all the values for id|title|author|price|imageurl")
		updateCmd.PrintDefaults()
		os.Exit(1)
	} else {

		book, _ := getBookById(*id)
		if (Book{}) == book {
			fmt.Println("Book Not found")
		} else {
			books := getBooks()

			for i, book := range books {
				if book.Id == *id {
					books[i] = Book{
						Id:       *id,
						Title:    *title,
						Author:   *author,
						Price:    *price,
						ImageUrl: *imageurl,
					}
				}
			}
			saveBooks(books)
		}

	}
	return

}

func handleAdd(addCmd *flag.FlagSet, id *string, title *string, author *string, price *string, imageurl *string) {

	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *author == "" || *price == "" || *imageurl == "" {
		fmt.Println("Pass all the values for id|title|author|price|imageurl")
		addCmd.PrintDefaults()
		os.Exit(1)
	} else {
		book := Book{
			Id:       *id,
			Title:    *title,
			Author:   *author,
			Price:    *price,
			ImageUrl: *imageurl,
		}

		books := getBooks()
		books = append(books, book)

		saveBooks(books)

	}
	return

}

func handleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	fmt.Println("get all books")
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("Pass --all or id option")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		books := getBooks()

		fmt.Printf("Id \t Title \t Author \t Price \t ImageURL \n")

		for _, book := range books {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price, book.ImageUrl)
		}
	}

	if *id != "" {
		id := *id
		book, _ := getBookById(id)
		fmt.Printf("Id \t Title \t Author \t Price \t ImageURL \n")
		fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price, book.ImageUrl)
	}

	return
}
