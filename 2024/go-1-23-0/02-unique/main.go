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
	u1 := unique.Make(User{Name: "Mario", LastName: "Carrion"})
	g2 := unique.Make(User{Name: "Mario", LastName: "Carrion"})

	fmt.Println("same values?", u1 == g2)

	// Group does not satisfy comparable
	// g1 := unique.Make(Group{Name: "Mario",Users: []User{}})
}
