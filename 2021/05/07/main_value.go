package main

import (
	"context"
	"fmt"
	"log"
)

type jwt string

const auth jwt = "JWT"

func main() {
	ctx := context.WithValue(context.Background(), auth, "Bearer hi")

	//-

	bearer := ctx.Value(auth)
	str, ok := bearer.(string)
	if !ok {
		log.Fatalln("not a string")
	}

	fmt.Println("value:", str)
}
