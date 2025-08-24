package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// wg.Add(1) // RIGHT: Call Add before starting the goroutine
	go func() {
		wg.Add(1) // WRONG: WaitGroup.Add called from inside new goroutine
		defer wg.Done()

		fmt.Println("Sleeping for 1 second")
		time.Sleep(1 * time.Second)
	}()

	wg.Wait()
	fmt.Println("Done")
}
