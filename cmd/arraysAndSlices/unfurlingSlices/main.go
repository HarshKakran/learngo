package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 54}
	fmt.Println("The sum is: ", sum(xi...))
}

func sum(ii ...int) int {
	var res int

	for _, v := range ii {
		res += v
	}

	return res
}
