package main

import (
	"fmt"
	"strconv"
)

// Availbale in this main package
//Declare and assign multiple variables
var (
	myName       string            // Declation without assigning value
	myLocation   string = "London" // Declation with assigning value
	myAge        int    = 31
	myDcLocation        = "Bangalore" // Declation without type and with assigning value
	myCompany           = "Tismo"
)

// myJob := "Devops Lead" is not possible outside of function
var myJob string = "DevOps Lead"

// I is available in all packages [ upper case variable is available to all]
var I int = 0

func main() {
	fmt.Printf("Varibale examples\n")
	fmt.Printf("Global I is %v,%T\n", I, I)
	fmt.Printf("My Company is %v,%T\n", myCompany, myCompany)
	myCompany := "Infosys" // Overwrite the global var
	myName = "Akilan"      // Changing the value
	fmt.Printf("My Name is %v,%T\n", myName, myName)
	fmt.Printf("My Location is %v,%T\n", myLocation, myLocation)
	fmt.Printf("My Age is %v,%T\n", myAge, myAge)
	fmt.Printf("My Job is %v,%T\n", myJob, myJob)
	myAge = 32
	var myNewAge string
	// Type conversion
	myNewAge = strconv.Itoa(myAge) // Int to Ascii value or else it will print space
	fmt.Printf("My Age is %v,%T\n", myAge, myAge)
	fmt.Printf("My new Age is %v,%T\n", myNewAge, myNewAge)
	fmt.Printf("My Company is %v,%T\n", myCompany, myCompany)
	fmt.Printf("My Company Location is %v,%T\n", myDcLocation, myDcLocation)
}
