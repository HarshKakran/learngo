package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	count := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := count
			v++
			count = v
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(count)
}
