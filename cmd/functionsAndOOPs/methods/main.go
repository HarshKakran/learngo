package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) printInfo() {
	fmt.Println(p)
}

func main() {
	p1 := person{
		fname: "Harsh",
		lname: "Kakran",
		age:   21,
	}

	p1.printInfo()
}
