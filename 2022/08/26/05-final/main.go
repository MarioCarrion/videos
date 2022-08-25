package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var address string

	flag.StringVar(&address, "address", "http://0.0.0.0:8300", "Vault address to use")

	flag.Parse()

	//

	config, err := NewConfig(os.Getenv("VAULT_TOKEN"), address)
	if err != nil {
		log.Fatalln("NewConfig", err)
	}

	val, err := config.Get(context.Background(), "password")
	if err != nil {
		log.Fatalln("config.Get", err)
	}

	fmt.Println("password:", val)
}
