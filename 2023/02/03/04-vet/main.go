package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-02-01")) // WARNING: 2006-02-01 should be 2006-01-02
}
