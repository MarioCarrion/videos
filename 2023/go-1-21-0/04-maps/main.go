package main

import (
	"fmt"
	"maps"
)

func main() {
	clients := map[string]int{
		"mario": 99,
		"ruby":  89,
	}

	fmt.Println(`clients := map[string]int{"mario": 99, "ruby": 89}`)

	// Clone
	cloned := maps.Clone(clients)
	fmt.Printf("\tcloned = %v, clients = %v\n", cloned, clients)
	fmt.Printf("\tcloned = %p, clients = %p\n", &cloned, &clients)

	// Equal
	fmt.Println("\n\tmaps.Equal(cloned, clients)", maps.Equal(clients, clients))

	// Copy
	dest := map[string]int{"mario": 0, "other": -1}
	fmt.Println("\n", `dest := map[string]int{"mario": 0, "other": -1}`)

	maps.Copy(dest, clients)
	fmt.Println("\tmaps.Copy(clients, dest) =", dest)

	// DeleteFunc
	maps.DeleteFunc(dest, func(_ string, v int) bool {
		if v > 90 {
			return true
		}

		return false
	})

	fmt.Println("\tmaps.DeleteFunc(dest, v > 90) =", dest)
}
