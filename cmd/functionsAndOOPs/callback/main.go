package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doOp)

	x := doOp(41, 16, add)
	fmt.Println(x)

	y := doOp(41, 16, subtract)
	fmt.Println(y)
}

// callbacks  --> passing functions as an argument in a function

func doOp(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
