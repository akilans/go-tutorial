package main

import "fmt"

var friends = [3]string{"Akilan", "Alex", "Jegan"}
var numbers = [...]int{1, 2, 3} // no need to specify array length. [...]

var letterMatrix [3][3]string = [3][3]string{
	[3]string{"A", "B", "C"},
	[3]string{"D", "E", "F"},
	[3]string{"G", "H", "I"},
}

var anotherLetterMatrix [2][3]string = [2][3]string{
	[3]string{"A", "B", "C"},
	[3]string{"D", "E", "F"},
}

var copiedFriends = friends // Exact copy, Change here does no impact in friends
var refFriends = &friends   //Pointer, Change here does impact in friends

// Slice Example
var sliceNumbers = []int{5, 10, 15}          // no need mention the length
var anotherSliceNumbers = make([]int, 3, 10) // type,length,capacity

func main() {
	fmt.Println("Arrays and Slice Example")
	fmt.Printf("Numbers list %v, %T \n", numbers, numbers)
	fmt.Printf("Friends list %v, %T \n", friends, friends) //Friends list [Akilan Alex Jegan], [3]string
	fmt.Printf("1st in Friends list %v, %T \n", friends[0], friends[0])
	numbers := [...]int{4, 5, 6}
	fmt.Printf("Numbers list %v, %T \n", numbers, numbers)
	//Letter matrix [[A B C] [D E F] [G H I]], [3][3]string
	fmt.Printf("Letter matrix %v, %T \n", letterMatrix, letterMatrix)
	// Another Letter matrix [[A B C] [D E F]], [2][3]string
	fmt.Printf("Another Letter matrix %v, %T \n", anotherLetterMatrix, anotherLetterMatrix)

	// Letter matrix[1][1] E, string
	fmt.Printf("Letter matrix[1][1] %v, %T \n", letterMatrix[1][1], letterMatrix[1][1])
	fmt.Printf("Friends list %v, %T \n", copiedFriends, copiedFriends)
	refFriends[0] = "Inba"
	// Friends list [Inba Alex Jegan], [3]string
	fmt.Printf("Friends list %v, %T \n", friends, friends)

	fmt.Printf("Total Friends %v \n", len(friends)) // 3

	// slice
	fmt.Printf("Total sliceNumbers %v \n", sliceNumbers)      //  [5 10 15]
	fmt.Printf("Total sliceNumbers %v \n", len(sliceNumbers)) // 3
	fmt.Printf("Total sliceNumbers %v \n", cap(sliceNumbers)) // 3

	fmt.Printf("Total anotherSliceNumbers %v \n", anotherSliceNumbers)      //  [0 0 0]
	fmt.Printf("Total anotherSliceNumbers %v \n", len(anotherSliceNumbers)) // 3
	fmt.Printf("Total anotherSliceNumbers %v \n", cap(anotherSliceNumbers)) // 10

	anotherSliceNumbers = sliceNumbers
	fmt.Printf("Total anotherSliceNumbers %v \n", anotherSliceNumbers)      //[5 10 15]
	fmt.Printf("Total anotherSliceNumbers %v \n", len(anotherSliceNumbers)) // 3
	fmt.Printf("Total anotherSliceNumbers %v \n", cap(anotherSliceNumbers)) // 3

	anotherNewSliceNumbers := append(anotherSliceNumbers[1:], []int{20, 25, 30}...)
	// [10 15 20 25 30] // remove 5
	fmt.Printf("Total anotherNewSliceNumbers %v \n", anotherNewSliceNumbers)
}
