package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

// CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo ./main.go

//go:embed LICENSE
var license string

func main() {
	//- Embedded Files

	fmt.Println(license)

	//- Deprecation of io/ioutil

	f, err := os.Open("LICENSE")
	if err != nil {
		log.Fatalln("Couldn't open file", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("reading all", err)
	}

	fmt.Printf("%s", b)

	//- os/signal.NotifyContext

	fmt.Println("waiting for signal")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	select {
	case <-time.After(time.Second):
		fmt.Println("exited after timer, missed signal")
	case <-ctx.Done():
		fmt.Println("signal received:", ctx.Err())
		stop()
	}

	fmt.Println("exiting")
}
