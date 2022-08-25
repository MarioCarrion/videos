package main

import (
	"flag"
	"fmt"
)

func main() {
	var password string

	flag.StringVar(&password, "password", "default", "password to use for connecting to the database")

	flag.Parse()

	fmt.Println("password:", password)
}
