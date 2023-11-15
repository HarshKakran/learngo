package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

// Any struct with the follwoing method is a stringer interface
// This defines how the TYPE will be printed
func (b book) String() string {
	return fmt.Sprint("This is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

// Wrapper function
func logInfo(s fmt.Stringer) {
	log.Println("This a log from the wrapper function", s.String())
}

func main() {
	b := book{
		title: "Ramayan - Unravelled",
	}

	var c count = 42

	logInfo(b)
	fmt.Println(c)
}
