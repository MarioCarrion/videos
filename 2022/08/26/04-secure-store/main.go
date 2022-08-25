package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	var address string

	flag.StringVar(&address, "address", "http://0.0.0.0:8300", "Vault address to use")

	flag.Parse()

	//-

	config := &vault.Config{
		Address: address,
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalln("vault.NewClient", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	logicalClient := client.Logical()

	secret, err := logicalClient.ReadWithContext(context.Background(), "/secret/data/credentials")
	if err != nil {
		log.Fatalln("client.ReadWithContext", err)
	}

	if secret == nil {
		log.Fatalln("secret is nil")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalln("data is missing")
	}

	val, ok := data["password"]
	if !ok {
		log.Fatalln("password is missing")
	}

	fmt.Println("password:", val)
}
