package main

import "fmt"

func main() {
	// send-only channel, can only send value
	sc := make(chan<- int)
	sc <- 42
	// fmt.Println(<-sc) // won;t work

	// receive-only channel, can only receive value
	rc := make(<-chan int)
	// rc <- 42 // won't work
	fmt.Println(<-rc)
}
