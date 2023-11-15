package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type foo int

func main() {
	var a foo = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	p1 := person{
		fname: "Harsh",
		lname: "Kakran",
		age:   21,
	}

	fmt.Printf("%#v", p1)
}
