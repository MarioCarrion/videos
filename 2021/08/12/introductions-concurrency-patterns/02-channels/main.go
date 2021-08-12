package main

import "fmt"

// XXX: This program panics because there is no goroutine
// outside of `main` interacting with the `ch` channel:
//
// fatal error: all goroutines are asleep - deadlock!

func main() {
	var ch chan int
	ch = make(chan int)
	// Two lines above can be also replaced with
	// ch := make(chan int)

	ch <- 10

	v := <-ch
	fmt.Println("received", v)
}
