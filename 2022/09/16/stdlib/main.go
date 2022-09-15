package main

import (
	"fmt"
	"net/url"
)

func main() {
	// https://pkg.go.dev/vuln/GO-2022-0988

	fmt.Println(url.JoinPath("https://go.dev", "../x"))  // https://go.dev/../x
	fmt.Println(url.JoinPath("https://go.dev/", "../x")) // https://go.dev/x
}
