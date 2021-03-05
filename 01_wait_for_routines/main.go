package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const n = 1000
	var wg = sync.WaitGroup{}
	//initialize wait-group
	wg.Add(n)
	// start n goroutines
	for i := 0; i < n; i++ {
		go routine(&wg)
	}
	// wait for all goroutines to finish their work
	wg.Wait()
	// print the number of active goroutines,
	// it is expected to get 1 active goroutine ( main )
	fmt.Printf("number of active go-routines is %d", runtime.NumGoroutine())
}

//routine is a function with simple logic and a wait-group countdown
func routine(wg *sync.WaitGroup) {
	// business logic
	i := 0
	i += i
	// wait-group countdown
	wg.Done()
}
