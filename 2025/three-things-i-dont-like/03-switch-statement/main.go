package main

import (
	"errors"
	"fmt"
)

func main() {
	Example1()
}

//-

func Example1() {
	value := 10
	switch value {
	case 2:
		fmt.Println("It's 2")
	case 10:
		fmt.Println("It's 10")
	}
}

//-

func Example2() {
	value := 1
	if value == 1 {
		fmt.Println("It's 1")
	} else if value == 2 {
		fmt.Println("It's 2")
	} else if value == 3 {
		fmt.Println("It's 3")
	} else {
		fmt.Println("It's something else")
	}
}

//-

func Example3() {
	str, err := Value("hello")

	switch {
	case errors.Is(err, ErrEmpty):
		fmt.Println("empty string")
	case errors.Is(err, ErrTooLong):
		fmt.Println("too long string")
	case str == "hello":
		fmt.Println("said hello")
	case str == "world":
		fmt.Println("said world")
	case err != nil:
		fmt.Println("other error", err)
	default:
		fmt.Println("value was", str)
	}
}

var ErrEmpty = errors.New("empty")
var ErrTooLong = errors.New("too long")

func Value(str string) (string, error) {
	if len(str) == 0 {
		return "", ErrEmpty
	}
	if len(str) > 5 {
		return "", ErrTooLong
	}

	if str == "hello" {
		return "world", nil
	}
	if str == "world" {
		return "hello", nil
	}

	return "hi", nil
}
