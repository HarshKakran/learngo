package main

import (
	"fmt"
)

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nTHe memory location of the thing2 array is : %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func squarePt(thing2 *[5]float64) *[5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The valeu p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	*p = 10
	// p now refers to i
	p = &i
	*p = 1

	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is %v", i)

	// Pointers and functions
	var thing1 = [5]float64{1, 2, 3, 4}
	fmt.Printf("\n Memory location for thing1 is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe reuslt is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}
