package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	fmt.Println("waiting")

	<-c

	fmt.Println("received")
}
