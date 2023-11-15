package main

import "fmt"

func main() {
	defer foo()
	bar()
}

// func(r receiver)indentifier(parameters)(return(s)) {<code>}
func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

/*
"defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because
- The surrounding fuction executed returns a statement
- reached the end of its function body
- or because the corresponding go routine is panicking
*/
