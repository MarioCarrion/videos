package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://url.xyz/?arg1=one;arg2=two&arg3=three")

	fmt.Println("err", err)
	fmt.Printf("%+v\n", u.Query())

	// Use https://pkg.go.dev/net/http#AllowQuerySemicolons to support previous behavior
}
