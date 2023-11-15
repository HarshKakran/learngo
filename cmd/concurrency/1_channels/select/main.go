package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i&1 != 0 {
			o <- i
		} else {
			e <- i
		}
	}
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel: ", v)
		case v := <-o:
			fmt.Println("from the odd channel: ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok ", i)
				return
			} else {
				fmt.Println("from comma ok ", i)
			}
		}
	}
}
