package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	/*
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
			close(c)
		}()
	*/
	go foo(c)

	//receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
