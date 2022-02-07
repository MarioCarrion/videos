package main

import (
	"dont-return-interfaces/producer1"
	"dont-return-interfaces/producer2"
	"fmt"
)

type Stringer interface {
	String() string
}

func main() {
	p1 := producer1.New()
	fmt.Println("name", p1.Produce())

	p2 := producer2.New()
	fmt.Println("car", p2.Produce())
}
