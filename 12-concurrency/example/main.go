// Routiens & Channels
// Normal execution steps executes one by one. Wait for previous one to finish
// Routines & Channels are came to solve that issue
// Create child routies inside main routie [ main() method ] by "go" keyword
// Make the channel variable with return type
// Get content from the routine & push into channels
// Main routine triggers the child routine & goes on sleep mode
// once child routines finish main routine stops

package main

import (
	"fmt"
	"net/http"
)

func main() {

	websites := []string{
		"http://facebook.com",
		"http://google.com",
		"http://tamilrockers.com",
		"http://twitter.com",
	}

	fmt.Println("Without Routines & channels started")
	for _, link := range websites {
		testLinkWithOutChannel(link)
	}

	c := make(chan string)

	fmt.Println("With Routines & channels started")
	for _, link := range websites {
		go testLink(link, c)
	}

	//fmt.Println(<-c) // Child routines executes first comes here
	//fmt.Println(<-c) // Child routines executes second comes here

	// All child routines execution process
	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}

}

func testLink(l string, c chan string) {

	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, " is down!!!")
		fmt.Println("Error : ", err)
		c <- "Server is down!!!"
		return
	}

	fmt.Println(l, " is up!!!")
	c <- "Server is up!!!"
}

func testLinkWithOutChannel(l string) {

	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, " is down!!!")
		fmt.Println("Error : ", err)
	}

	fmt.Println(l, " is up!!!")
}
