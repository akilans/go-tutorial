package main

import "fmt"

const name string = "Akilan"

const (
	myName string = "Akilan"
	myAge  int    = 32
)

const (
	index0 = iota // iota is a special type. It will increment the next const
	index1
	index2
	index3
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KB bytewise
	KB = 1 << (10 * iota) // 1024 => 10000000000 // left shift 10
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("Hello %v\n", name)
	const name string = "Inba" // You can redeclare const within block
	fmt.Printf("Hello %v\n", name)
	fmt.Printf("Hello %v\n", myName)
	fmt.Printf("Your age is %v\n", myAge)
	fmt.Printf("First index of my name is %v\n", string(myName[index0]))
	fmt.Printf("Second index of my name is %v\n", string(myName[index1]))
	fmt.Printf("Third index of my name is %v\n", string(myName[index2]))
	fmt.Printf("Fourth index of my name is %v\n", string(myName[index3]))

	fmt.Printf("%v MB , %T %v\n", MB, MB, MB/KB)
	fmt.Printf("%v GB , %T %v\n", GB, GB, GB/MB)

	fileSize := 40000000000.00
	fmt.Printf("%.2fGB\n", fileSize/GB) // 37.25GB
}
