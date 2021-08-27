package main

import "fmt"

type CustomError int

func (_ CustomError) Error() string {
	return "error message"
}

func (_ CustomError) Is(err error) bool {
	return false
}

func (_ CustomError) As(target interface{}) bool {
	return false
}

func (_ CustomError) Unwrap() error {
	return nil
}

func main() {
	var err error

	err = CustomError(10)

	fmt.Printf("Error value: %v", err)
}
