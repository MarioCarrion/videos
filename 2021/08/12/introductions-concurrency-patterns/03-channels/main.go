package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(time.Now(), "taking a nap")

		time.Sleep(2 * time.Second)

		ch <- "hello"
	}()

	fmt.Println(time.Now(), "waiting for message")

	v := <-ch

	fmt.Println(time.Now(), "received", v)
}
