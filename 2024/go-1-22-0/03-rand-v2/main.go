package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	msgs := []string{
		"one",
		"two",
		"three",
		"four",
	}

	fmt.Println("number is:", msgs[rand.IntN(len(msgs))])

	// IntXYZ are idiomatic, for example:
	// math/rand/Int31 -> math/rand/v2/Int32

	fmt.Println(rand.Int32())
	fmt.Println(rand.Int64())

	// rand.N(5*time.Minute)

	fmt.Println(rand.N(10 * time.Second))
}
