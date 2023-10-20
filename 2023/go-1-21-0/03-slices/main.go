package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices.Index
	indexNumbers := []int{100, 2, 99, 2}
	fmt.Println(`indexNumbers := []int{100, 2, 99, 2}`)
	fmt.Println("\tslices.Index(indexNumbers, 1) =", slices.Index(indexNumbers, 1))       // Returns -1, not found
	fmt.Println("\tslices.Index(indexNumbers, 2) =", slices.Index(indexNumbers, 2), "\n") // Returns 1

	// slices.Sort
	sortNumbers := []int64{100, 2, 99, 2}
	slices.Sort(sortNumbers)
	fmt.Println(`sortNumbers := []int64{100, 2, 99, 2}`)
	fmt.Println("\tslices.Sort(sortNumbers) =", sortNumbers, "\n")

	// slices.Min + slices.Max
	minMaxNumbers := []int{100, 2, 99, 2}
	fmt.Println(`minMaxNumbers := []int{100, 2, 99, 2}`)
	fmt.Println("\tslices.Max(minMaxNumbers)", slices.Max(minMaxNumbers))
	fmt.Println("\tslices.Min(minMaxNumbers)", slices.Min(minMaxNumbers))
}
