package main

import (
	"fmt"
	"unique"
)

// `comparable`
type User struct {
	Name     string
	LastName string
}

// Not `comparable`
type Group struct {
	Name  string
	Users []User
}

func main() {
	h1 := unique.Make(User{Name: "Mario", LastName: "Carrion"})
	h2 := unique.Make(User{Name: "Mario", LastName: "Carrion"})

	fmt.Println("same values?", h1 == h2)
	fmt.Printf("addresses: %v - %v\n", h1, h2)

	// Group does not satisfy comparable
	// g1 := unique.Make(Group{Name: "Mario",Users: []User{}})
}
