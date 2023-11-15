package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	ownerInfo2 owner
}

// struct method
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	// method signature
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Kakran"}, owner{"Harsh"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo2.name, myEngine.name)
	fmt.Printf("Total miles left in the tank: %v", myEngine.milesLeft())

	canMakeIt(myEngine, 50)

	fmt.Println(myEngine.milesLeft())

}
