package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
	fname string
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
			age:   32,
		},
		fname: "Ethan Hawk",
		ltk:   true,
	}

	fmt.Printf("%T \t %v\n", sa1, sa1)
	fmt.Println(sa1.age, sa1.fname, sa1.ltk, sa1.person.fname)
}
