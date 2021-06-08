package main

import "fmt"

func main() {
	number := 4
	fmt.Println(square(number))
}

func square(number int) int {
	return number + number
}
