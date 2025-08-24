package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("runtime.NumCPU:", runtime.NumCPU())
}
