package main

import "fmt"

// Grade system
var mark int = 101

var akilanDetails = map[string]string{
	"name":     "Akilan",
	"location": "London",
}

func main() {

	var keyName string = "name"

	// Key exists check
	if val, ok := akilanDetails[keyName]; ok {
		fmt.Println(ok)
		fmt.Println("Key " + keyName + " Exists. & the value is " + val)
	}

	if val, isExists := akilanDetails["age"]; isExists {
		fmt.Println(val)
		fmt.Println(isExists)
	} else {
		fmt.Println(isExists)
	}

	if mark < 0 || mark > 100 {
		fmt.Println("Please enter mark between 0 - 100!!!")
	} else if mark >= 0 && mark < 35 {
		fmt.Println("Failed!!!")
	} else if mark >= 35 && mark <= 50 {
		fmt.Println("Average!!!")
	} else if mark > 50 && mark <= 75 {
		fmt.Println("Good!!!")
	} else if mark > 75 && mark <= 90 {
		fmt.Println("Very Good!!!")
	} else if mark > 90 && mark <= 100 {
		fmt.Println("Awesome!!!")
	} else {
		fmt.Println("No Idea!!!")
	}

	switch {
	case mark < 0 || mark > 100:
		fmt.Println("Please enter mark between 0 - 100!!!")
	case mark >= 0 && mark < 35:
		fmt.Println("Failed!!!")
	case mark >= 35 && mark <= 50:
		fmt.Println("Average!!!")
	case mark > 50 && mark <= 75:
		fmt.Println("Good!!!")
	case mark > 75 && mark <= 90:
		fmt.Println("Very Good!!!")
	case mark > 90 && mark <= 100:
		fmt.Println("Awesome!!!")
	default:
		fmt.Println("No Idea!!!")

	}

}
