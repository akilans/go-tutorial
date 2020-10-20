package main

import (
	"fmt"
	"os"
)

func main() {

	data := "Akilan is a DevOps Lead.\nHe is from India. \nHe is working in London"

	newFile := createFile("akilan.txt")
	defer closeFile(newFile) // It will get call at last
	writeFile(newFile, data)
}

func writeFile(fileName *os.File, data string) {
	fmt.Println("Writing a file...")
	fmt.Fprintf(fileName, data)
}

func createFile(p string) *os.File {
	fmt.Println("Creating a file...")
	res, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return res
}

func closeFile(filePath *os.File) {
	fmt.Println("Closing a file...")
	err := filePath.Close()
	if err != nil {
		panic(err)
	}
}
