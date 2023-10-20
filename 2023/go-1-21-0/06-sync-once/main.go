package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sync.OnceValue (sync.OnceValues works the same)

	onceString := sync.OnceValue[string](func() string {
		time.Sleep(time.Second)
		return "sync.OnceValue: computed once!"
	})

	fmt.Println(time.Now(), onceString())
	fmt.Println(time.Now(), onceString())
	fmt.Println(time.Now(), onceString())

	// sync.OnceFunc
	fmt.Println("")

	onceFunc := sync.OnceFunc(func() {
		time.Sleep(time.Second)
		fmt.Println("sync.OnceFunc: computed once!")
	})

	fmt.Println(time.Now())
	onceFunc()
	fmt.Println(time.Now())
	onceFunc()
	fmt.Println(time.Now())
	onceFunc()
}
