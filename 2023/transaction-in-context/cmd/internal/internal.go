package conn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

func NewConnString() string {
	// XXX: Do not hardcode configuration values
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("user", "password"),
		Host:   "localhost",
		Path:   "dbname",
	}

	return dsn.String()
}

func Print(val any) {
	res, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", res)
}
