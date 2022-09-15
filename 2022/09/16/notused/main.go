package main

import (
	"gopkg.in/yaml.v3"
)

func main() {
	type T struct {
		F int `yaml:"a,omitempty"`
		B int
	}

	yaml.Marshal(&T{B: 2}) // Returns "b: 2\n"
}
