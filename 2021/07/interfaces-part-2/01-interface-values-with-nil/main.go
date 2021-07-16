package main

import "fmt"

type Interface interface {
	Method()
}

type String struct {
	Value string
}

func (s *String) Method() {
	fmt.Printf("s= %v\n", s)
}

func main() {
	var str *String
	fmt.Printf("str= %v\n", str) // <nil>
	str.Method()                 // <nil>

	(*String).Method(str) // Syntatic sugar

	var iface Interface = str
	iface.Method() // <nil>
}
