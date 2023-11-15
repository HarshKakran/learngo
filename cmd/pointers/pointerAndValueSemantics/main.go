package main

import "fmt"

// Value semantics
func addOne(a int) int {
	return a + 1
}

// Pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {
	// value semantics
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Printf("%v\n\n", a)

	// pointer semantics
	b := 1
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
}
