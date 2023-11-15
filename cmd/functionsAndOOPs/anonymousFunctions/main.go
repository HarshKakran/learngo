package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Printf("This is an anonymous func with %v\n", s)
	}("parameters")
}

func foo() {
	fmt.Println("Foo Ran")
}

// named function
// func(r receiver)identifier(p parameter(s))(r return(s)){code}

// anonymous functino
// func (p parameter(s))(r return(s)){ code }
