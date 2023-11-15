package main

import "fmt"

func main() {
	/* Won't Work
	c := make(chan int)
	c <- 42
	*/

	// One way of successful pass
	// c := make(chan int)
	// go func() {
	// 	c <- 42
	// }()

	// 2nd way of successful pass/buffer
	// c := make(chan int, 1)
	// c <- 42

	// Unsuccessful buufer
	c := make(chan int, 1)
	c <- 42
	c <- 43

	fmt.Println(<-c)
}
