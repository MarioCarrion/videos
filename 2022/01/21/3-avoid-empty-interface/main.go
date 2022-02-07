package main

import "fmt"

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Mario"
	person["age"] = 99

	for _, val := range person {
		switch x := val.(type) {
		case string:
			Do(x)
		}
	}
}

func Do(v string) {
	fmt.Println("this is a string", v)
}
