package example

import "fmt"

type Value struct {
	MustBeTrue bool
	Message    string
	Priority   int
}

// XXX: Do not do this
func ParseValue(value Value) (string, error) {
	if !value.MustBeTrue {
		return "", fmt.Errorf("must be true")
	} else {
		if value.Message == "" {
			return "", fmt.Errorf("value must be set")
		} else {
			if value.Priority == 0 {
				return "", fmt.Errorf("priority must be set")
			}

			return fmt.Sprintf("%d %s", value.Priority, value.Message), nil
		}
	}
}
