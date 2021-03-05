package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 1000

	counter := uint32(0)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	//initialize wait-group
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done() // wait-group decrement
			// mutex lock
			mutex.Lock()
			// defer mutex unlock
			defer mutex.Unlock()
			counter++
		}(&wg)
	}
	// wait for all goroutines to finish their work
	wg.Wait()
	// print counter
	fmt.Println(counter)
}
