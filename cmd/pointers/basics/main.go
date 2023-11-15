package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	s := "Harsh"
	fmt.Printf("%v\t%T\n", &s, &s)

	// Pointer declaration
	// y := &x
	var y *int = &x
	fmt.Println(y)
	// Dereferencing a pointer --> *y
	fmt.Println(*y, *&x)

}
