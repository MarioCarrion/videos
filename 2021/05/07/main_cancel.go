package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	run := func(ctx context.Context) {
		n := 1
		for {
			select {
			case <-ctx.Done(): // 2. "ctx" is cancelled, we close "ch"
				fmt.Println("exiting")
				close(ch)
				return // returning not to leak the goroutine
			default:
				time.Sleep(time.Millisecond * 300)
				fmt.Println(n)
				n++
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("goodbye")
		cancel() // 1. cancels "ctx"
	}()

	go run(ctx)

	fmt.Println("waiting to cancel...")

	<-ch // 3. "ch" is closed, we exit

	fmt.Println("bye")
}
