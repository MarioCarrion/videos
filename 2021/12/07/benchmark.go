package benchmark

import "strings"

func Concat(a, b string) string {
	// Ineffecient implementation using fmt
	// return fmt.Sprintf("%s%s", a, b)

	// Efficient implementation starts here
	builder := strings.Builder{}

	builder.WriteString(a)
	builder.WriteString(b)

	return builder.String()
}
