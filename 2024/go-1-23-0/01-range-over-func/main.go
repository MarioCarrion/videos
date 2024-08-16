package main

import (
	"fmt"
	"time"
)

// func(func() bool)
func Iter1() func(func() bool) {
	return func(yield func() bool) {
		var i int

		for {
			if !yield() {
				return
			}
			i++

			fmt.Println(i, time.Now())
		}
	}
}

// func(func(K) bool)
func Iter2() func(func(int) bool) {
	return func(yield func(int) bool) {
		var v int

		for {
			if !yield(v) {
				return
			}
			v += 10
		}
	}
}

// func(func(K, V) bool)
func Iter3() func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		var v int

		for {
			if !yield(v, v+10) {
				return
			}
			v++
		}
	}
}

func MultiplyBy2(vals ...int) []int {
	res := make([]int, len(vals))
	for i, v := range vals {
		res[i] = v * 2
	}
	return res
}

func MultiplyBy2Iter(vals ...int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for i, v := range vals {
			if !yield(i, v*2) {
				return
			}
		}
	}
}

func main() {
	var i = 0
	for range Iter1() { // can't "for _ = range Basic()"
		if i == 5 {
			break
		}
		i++
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("---")

	for v := range Iter2() {
		if v == 50 {
			break
		}

		fmt.Println(v, time.Now())
	}

	fmt.Println("---")

	for i, v := range Iter3() {
		if i == 5 {
			break
		}

		fmt.Println(i, v, time.Now())
	}

	//-

	fmt.Println("---")

	for i, v := range MultiplyBy2(1, 2, 3, 4, 5, 6) {
		fmt.Println(i, v)
	}

	fmt.Println("---")

	for i, v := range MultiplyBy2Iter(1, 2, 3, 4, 5, 6) {
		fmt.Println(i, v)
	}
}
