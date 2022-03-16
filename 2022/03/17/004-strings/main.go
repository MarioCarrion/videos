package main

import (
	"fmt"
	"strings"
)

// It can replace and simplify many common uses of Index, IndexByte, IndexRune, and SplitN.

func main() {
	cut := func(sep string) {
		s := "hello|world"
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("cut(%q, %q): %q, %q, %v\n", s, sep, before, after, found)
	}

	cut("|")
	cut("hello")
	cut("nothing")
}
