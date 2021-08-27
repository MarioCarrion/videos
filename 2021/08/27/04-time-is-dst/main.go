package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("DST?", time.Now().IsDST())
}
