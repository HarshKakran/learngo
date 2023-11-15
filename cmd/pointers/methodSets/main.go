package main

import "fmt"

type dog struct {
	fname string
}

func (d dog) walk() {
	fmt.Println("My name is ", d.fname, "and I am walking,")
}

func (d *dog) run() {
	d.fname = "Rover"
	fmt.Println("My name is ", d.fname, "and I am running,")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) --> Error
	youngRun(&d1)

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)
}
