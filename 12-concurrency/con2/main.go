package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // 2 go routines

	go func() { // 1st one - Go routines anonymous function
		process("ordering....")
		wg.Done()
	}()

	go func() { // 2nd one - Go routines anonymous function
		process("paying....")
		wg.Done()
	}()
	wg.Wait()

}

func process(action string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 1)
		fmt.Println(i, action)
	}

}
