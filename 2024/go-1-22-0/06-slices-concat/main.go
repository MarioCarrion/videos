package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []string{"one", "two"}
	s2 := []string{"three"}
	s3 := []string{"four"}

	res := slices.Concat[[]string](s1, s2, s3)

	fmt.Printf("slices: %#v\n", res)
}
