package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:          ", runtime.GOOS)
	fmt.Println("ARCH:        ", runtime.GOARCH)
	fmt.Println("CPUs:        ", runtime.NumCPU())
	fmt.Println("GO Routines: ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("\nCPUs:        ", runtime.NumCPU())
	fmt.Println("GO Routines: ", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar: ", i)
	}
}
