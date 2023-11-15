package main

import "fmt"

type person struct {
	fname string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am ", p.fname)
}
func (sa secretAgent) speak() {
	fmt.Printf("I am %v and ", sa.fname)
	if sa.ltk {
		fmt.Println("I have license to kill.")
	} else {
		fmt.Println("I don't have a license to kill.")
	}
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			fname: "James",
		},
		ltk: true,
	}

	p1 := person{
		fname: "Jenny",
	}

	/*
		sa1.speak()
		p1.speak()
	*/
	saySomething(sa1)
	saySomething(p1)
}

/*
AN interface in GO defines a set of method signatures
Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE.
*/

/*
package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}

func (s square) area() float64 {
	return s.l * s.w
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

type shape interface {
	area() float64
}

func calArea(s shape) float64 {
	return s.area()
}

func main() {
	sq := square{4, 4}
	c := circle{4}

	fmt.Println(calArea(c))
	fmt.Println(calArea(sq))
}

*/
