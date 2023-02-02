package main

import (
	"errors"
	"fmt"
	"io"
)

var ErrOne = fmt.Errorf("one")
var ErrTwo = fmt.Errorf("two")

func main() {
	// Using errors.Join
	join := errors.Join(ErrOne, ErrTwo)
	fmt.Printf("%+v\n", join)

	if errors.Is(join, ErrOne) {
		fmt.Println("One!")
	}

	if errors.Is(join, ErrTwo) {
		fmt.Println("Two!")
	}

	// Using fmt.Errorf with multiple %w
	efmt := fmt.Errorf("%w, %w", errors.New("fmt"), io.EOF)

	fmt.Printf("%+v\n", efmt)

	if errors.Is(efmt, io.EOF) {
		fmt.Println("End of File")
	}
}
