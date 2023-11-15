package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("This function is assigned to variable X")
	}

	y := func(s string) {
		fmt.Println("This is a ", s, " function, assigned to the variable Y")

	}
	x()
	y("parameterized")
}
