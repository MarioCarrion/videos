package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	fmt.Printf("numbers = %v\n", numbers)

	// maps.Values
	fmt.Printf("\nslices.Sorted(maps.Values(numbers))\n")
	sortedValues := slices.Sorted(maps.Values(numbers))
	fmt.Printf("\t%v\n", sortedValues)

	// maps.Keys
	fmt.Printf("\nslices.Sorted(maps.Keys(numbers))\n")
	sortedKeys := slices.Sorted(maps.Keys(numbers))
	fmt.Printf("\t%q\n", sortedKeys)
}
