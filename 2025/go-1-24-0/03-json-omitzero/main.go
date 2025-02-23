package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Example struct {
	String       string    `json:"string,omitempty"`
	KeyValue     KeyValue  `json:"keyValue,omitzero"`
	OmitKeyValue KeyValue  `json:"omitKeyValue,omitzero"`
	ZeroValue1   ZeroValue `json:"zeroValue1,omitzero,omitempty"`
	ZeroValue2   ZeroValue `json:"zeroValue2,omitzero,omitempty"`
}

type KeyValue struct {
	Key   string
	Value string
}

func (z KeyValue) IsZero() bool {
	return z.Key == "" && z.Value == ""
}

type ZeroValue int64

func (z ZeroValue) IsZero() bool {
	return int64(z) < 10
}

func main() {
	example := Example{
		KeyValue: KeyValue{
			Key:   "key",
			Value: "value",
		},
		ZeroValue1: 5, // Not marshaled values IsZero returns true
		ZeroValue2: 0, // Not marshaled because it's an actual _zero value_
	}

	res, err := json.MarshalIndent(example, "", "    ")
	if err != nil {
		log.Fatalln("Couldn't marshal JSON:", err)
	}

	fmt.Println(string(res))
}
