package main

import (
	"fmt"
	"reflect"
)

var friends = []string{"Akilan", "Jegan", "Alex"}

var letterMatrix = [3][3]string{
	[3]string{"A", "B", "C"},
	[3]string{"D", "E", "F"},
	[3]string{"G", "H", "I"},
}

var akilanDetails = map[string]string{
	"name":     "Akilan",
	"location": "London",
}

func main() {

	fmt.Println(reflect.TypeOf(friends))

	// Loop 1
	for i := 0; i < len(friends); i++ {
		fmt.Println(friends[i])
	}

	// Loop 2
	j := 0
	for j < len(friends) {
		fmt.Println(friends[j])
		j++
	}

	//
	for k := 0; k < len(letterMatrix); k++ {
		for l := 0; l < len(letterMatrix[k]); l++ {
			fmt.Println(letterMatrix[k][l])
		}
	}

	// Looping maps
	for key, value := range akilanDetails {
		fmt.Println(key + " ===> " + value)
	}

}
