package main

import "fmt"

func main() {
	// `min`

	fmt.Println("min(1,0,-1) =", min(1, 0, -1))                     // Outputs: -1
	fmt.Println(`min("c","cat","a") =`, min("c", "cat", "a"), "\n") // Outputs: "a"

	// `max`

	fmt.Println("max(10,100,3) =", max(10, 100, 3))                         // Outputs: 100
	fmt.Println(`max("dat","xaz","xyz") =`, max("dat", "xaz", "xyz"), "\n") // Outputs: "xyz"

	// `clear` using slice

	strs := make([]string, 2)
	strs[0] = "hello"
	strs[1] = "world"

	fmt.Printf("%#v - %p\n", strs, &strs) // []string{"hello", "world"}

	clear(strs)
	fmt.Printf("%#v - %p\n\n", strs, &strs) // []string{"", ""}

	// `clear` using map

	clients := make(map[string]int)
	clients["mario"] = 99
	clients["ruby"] = 89

	fmt.Printf("%#v - %p\n", clients, &clients)
	clear(clients)

	fmt.Printf("%#v -%p\n", clients, &clients)
}
