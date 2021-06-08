package main

import (
	"fmt"
	"time"
)

func main() {
	go process("ordering....")
	go process("paying....")

	fmt.Scanln() // one way to stop the go routines

}

func process(action string) {
	for i := 0; true; i++ {
		time.Sleep(time.Second / 1)
		fmt.Println(action)
	}

}
