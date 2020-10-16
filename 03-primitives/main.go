package main

import "fmt"

func main() {
	fmt.Printf("Primitive Examples\n")
	a := 10
	b := 3
	// Int operation
	fmt.Printf("Addition of %v and %v is %v\n", a, b, (a + b))
	fmt.Printf("Subtraction of %v and %v is %v\n", a, b, (a - b))
	fmt.Printf("Multiplication of %v and %v is %v\n", a, b, (a * b))
	fmt.Printf("Remainder of %v and %v is %v\n", a, b, (a % b))

	// Boolean
	isGoGood := true
	c := a == b    // a != b so it is false
	fmt.Println(c) // false
	fmt.Printf("Is Go Good? %v\n", isGoGood)
	fmt.Printf("Is %v greater than %v is %v\n", a, b, (a >= b))

	// Bitwise operator
	// a =10 -> 1010
	// b =3 -> 0011
	fmt.Printf("AND operation of %v and %v is %v\n", a, b, (a & b))   // 2 -> 0010
	fmt.Printf("OR operation of %v and %v is %v\n", a, b, (a | b))    // 11 => 1011
	fmt.Printf("XOR operation of %v and %v is %v\n", a, b, (a ^ b))   // 9 => 1001
	fmt.Printf("XAND operation of %v and %v is %v\n", a, b, (a &^ b)) // 8 => 1000
	fmt.Printf("Left shift operation of %v  is %v\n", a, (a << 1))    // 20 => 10100
	fmt.Printf("Right shift operation of %v  is %v\n", a, (a >> 1))   // 5 => 0101

	// string
	introText := "Hi, I am Akilan"
	fmt.Printf("Intro text is %v\n", introText) // 72 => ascii value is H
	fmt.Printf("1st letter of Intro text is %v\n", introText[0])

	bytesOfIntro := []byte(introText)
	fmt.Printf("Bytes of Intro text is %v\n", bytesOfIntro)
}
