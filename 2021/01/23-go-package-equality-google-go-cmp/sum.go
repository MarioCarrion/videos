package operation

import "strings"

// SumSimple ...
func SumSimple(a, b int) int {
	return a + b
}

// Dollar ...
type Dollar struct {
	Superunit int
	Subunit   int
}

// SumDollar ...
func SumDollars(a, b Dollar) Dollar {
	const cents int = 100

	superunit := a.Superunit + b.Superunit
	subunit := a.Subunit + b.Subunit

	superunit += subunit / cents
	subunit %= cents

	return Dollar{
		Superunit: superunit,
		Subunit:   subunit,
	}
}

//-

// ConcatenateSlice ...
func ConcatenateSlice(a, b []int) []int {
	res := make([]int, len(a), len(a)+len(b))
	_ = copy(res, a)
	res = append(res, b...)
	return res
}

// Message ...
type Message string

// Equal ...
func (m Message) Equal(b Message) bool {
	return strings.ToLower(string(m)) == strings.ToLower(string(b))
}

// Alert ...
type Alert struct {
	Message Message
	code    int
}
