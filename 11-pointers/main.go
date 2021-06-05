package main

import "fmt"

func main() {
	myName := "Akilan"
	newName := &myName

	fmt.Println(newName)  // 0xc000010200
	fmt.Println(*newName) // Akilan

	akilan := Employee{
		name: "Akilan",
		age:  34,
	}

	inba := &akilan

	fmt.Println(akilan)
	fmt.Println(inba.name) // No need to mention *

}

type Employee struct {
	name string
	age  int
}
