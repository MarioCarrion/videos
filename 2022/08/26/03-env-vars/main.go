package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("password:", os.Getenv("PASSWORD"))
}
