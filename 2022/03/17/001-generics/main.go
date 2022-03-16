package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Sum[V constraints.Float | constraints.Integer](m ...V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	fmt.Println(Sum([]int64{1, 2, 3, 4}...))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(uint32(1), uint32(3)))
	fmt.Println(Sum(1.2, 2.3, 3.5, 4.6))
}
