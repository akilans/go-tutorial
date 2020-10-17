package main

import "fmt"

// Friends structure
type Friends struct {
	name      string
	age       int
	locations []string
}

// Map example => Key and value pair
var akilanDetails = map[string]string{
	"name":     "Akilan",
	"location": "London",
}

func main() {
	fmt.Println("Map & Structs Examples")
	// Access map values
	fmt.Printf("Name is %v and his location is %v\n", akilanDetails["name"], akilanDetails["location"])
	akilanDetails["role"] = "DevOps Lead"

	// Add new key values in map
	fmt.Printf("Role is %v\n", akilanDetails["role"])
	fmt.Printf("Akilan's detail is %v\n", akilanDetails)
	delete(akilanDetails, "role")
	fmt.Printf("Akilan's detail is %v\n", akilanDetails)

	alex := Friends{
		name: "Alex",
		age:  32,
		locations: []string{
			"Chennai",
			"Bangalore",
			"Pune",
		},
	}

	fmt.Printf("Alex's Details %v\n", alex)
	fmt.Printf("Alex's locations %v\n", alex.locations)

}
