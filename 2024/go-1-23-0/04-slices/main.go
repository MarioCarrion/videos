package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []string{"5", "4", "3", "2", "1"}
	fmt.Printf("numbers = %v\n", numbers)

	// slices.Chunk
	fmt.Printf("\nslices.Chunk(numbers, 2)\n")
	for c := range slices.Chunk(numbers, 2) {
		fmt.Printf("\t%s\n", c)
	}

	// slices.Collect
	fmt.Printf("\nslices.Collect(slices.Chunk(numbers, 2))\n")
	numbers1 := slices.Collect(slices.Chunk(numbers, 2))
	fmt.Printf("\t%s\n", numbers1)

	// slices.Sorted + slices.Values
	fmt.Printf("\nslices.Sorted(slices.Values(numbers))\n")
	sorted := slices.Sorted(slices.Values(numbers))
	fmt.Printf("\t%v\n", sorted)

	fmt.Printf("\nslices.Backward(numbers)\n")
	for i, v := range slices.Backward(numbers) {
		fmt.Printf("\ti: %d, v: %s\n", i, v)
	}
}
