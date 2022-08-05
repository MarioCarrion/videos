package main

import (
	"errors"
	"fmt"
)

func main() {
	var newErr = errors.New("sentinel")
	err := errors.New("foo")

	if errors.As(err, &newErr) {
		fmt.Println("error here!")
	}
}
