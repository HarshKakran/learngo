package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func main() {
	// anonymous struct
	p1 := struct {
		fname string
		lname string
		age   int
	}{
		fname: "Harsh",
		lname: "Kakran",
		age:   21,
	}

	// normal struct
	p2 := person{
		fname: "James",
		lname: "Bond",
		age:   32,
	}

	fmt.Printf("%T \t %v\n", p1, p1)
	fmt.Printf("%T \t %v\n", p2, p2)
}
