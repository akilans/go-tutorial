package main

import "fmt"

// Grade system
var mark int = 80

func main() {
	if mark < 0 || mark > 100 {
		fmt.Print("Please enter mark between 0 - 100 ")
	} else if mark >= 0 && mark < 35 {
		fmt.Print("Failed!!!")
	} else if mark >= 35 && mark <= 50 {
		fmt.Print("Average!!!")
	} else if mark > 50 && mark <= 75 {
		fmt.Print("Good!!!")
	} else if mark > 75 && mark <= 90 {
		fmt.Print("Very Good!!!")
	} else if mark > 90 && mark <= 100 {
		fmt.Print("Awesome!!!")
	} else {
		fmt.Print("No Idea!!!")
	}

	fmt.Print("\n")

}
