package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		panic("Some error happened")
		fmt.Println("It will not get printed")
	*/

	// Permisssion denied error
	res, err := os.Create("/etc/akilan.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
