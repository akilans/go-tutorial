package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go process("ordering....", out1)
	go process("Paying....", out2)

	for msg1 := range out1 {
		fmt.Println(msg1)
	}
	// print 5 ordering after that only it will go for paying
	/* 	ordering....
	   	ordering....
	   	ordering....
	   	ordering....
	   	ordering....
	   	Paying....
	   	Paying....
	   	Paying....
	   	Paying....
	   	Paying.... */

	for msg2 := range out2 {
		fmt.Println(msg2)
	}

}

func process(action string, out chan string) {
	defer close(out) // It will call last and best way to close channel
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 1)
		out <- action
	}

}
