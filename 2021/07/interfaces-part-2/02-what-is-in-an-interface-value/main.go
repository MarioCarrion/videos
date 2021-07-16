package main

import (
	"fmt"
)

type Interface interface {
	Method()
}

type String struct {
	Value string
}

func (s *String) Method() {}

type Integer int

func (i Integer) Method() {}

func main() {
	var iface Interface

	iface = &String{"hello world"}
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)

	iface = Integer(100)
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)
}
