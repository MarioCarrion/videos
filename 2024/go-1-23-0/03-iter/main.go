package main

import (
	"fmt"
	"iter"
	"time"
)

// iter.Seq[V any]     func(yield func(V) bool)    -- One value
// iter.Seq2[K, V any] func(yield func(K, V) bool) -- Two values: typically key-value or index-value pairs

// func(func(K, V) bool)
func Iter() func(func(int, int) bool) {
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

func main() {
	for i, v := range Iter() {
		if i == 5 {
			break
		}

		fmt.Println(i, v, time.Now())
	}

	//-

	next, stop := iter.Pull2(Iter())
	defer stop()

	for {
		i, v, ok := next()
		if !ok {
			break
		}

		if i == 5 {
			break
		}

		fmt.Println(i, v, time.Now())
	}
}
