package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 30, 4, 5, 6))
}

func sum(ii ...int) int {
	fmt.Printf("%T \t %v\n", ii, ii)

	var res int

	for _, v := range ii {
		res += v
	}

	return res
}

/*
When we create a function which can take unlimited arguments, this is known as fucntion with variadic parameters.
*/
