package main

import "fmt"

func main() {
	c := gen()

	for v := range c {
		fmt.Println(v)
	}
}

func gen() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}
		close(c)
	}()

	return c
}
