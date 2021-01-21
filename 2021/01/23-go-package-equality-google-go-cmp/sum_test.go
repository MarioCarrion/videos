package operation_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	operation "github.com/MarioCarrion/videos/2021/01/23-go-package-equality-google-go-cmp"
)

// PRESENT OMIT
func TestAbs(t *testing.T) {
	if actual := math.Abs(-1); actual != 1 {
		t.Fatalf("expected 1, actual %f", actual)
	}
}

// PRESENT END OMIT

func TestSumSimple(t *testing.T) {
	t.Parallel()

	type input struct {
		a int
		b int
	}

	tests := []struct {
		name     string
		input    input
		expected int
	}{
		{
			"OK: 1+2",
			input{1, 2},
			3,
		},
		{
			"OK: 2-2",
			input{2, -2},
			0,
		},
		// {
		// 	"ERROR: 1-2",
		// 	input{1, -2},
		// 	0,
		// },
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if actual := operation.SumSimple(test.input.a, test.input.b); actual != test.expected {
				t.Fatalf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

func TestSumDollars_PrimitiveTypes(t *testing.T) {
	t.Parallel()

	type input struct {
		a operation.Dollar
		b operation.Dollar
	}

	tests := []struct {
		name     string
		input    input
		expected operation.Dollar
	}{
		{
			"OK: Subunit does not exceed",
			input{
				operation.Dollar{
					Superunit: 2,
					Subunit:   20,
				},
				operation.Dollar{
					Superunit: 1,
					Subunit:   40,
				},
			},
			operation.Dollar{
				Superunit: 3,
				Subunit:   60,
			},
		},
		{
			"OK: Subunit exceeds",
			input{
				operation.Dollar{
					Superunit: 1,
					Subunit:   20,
				},
				operation.Dollar{
					Superunit: 1,
					Subunit:   85,
				},
			},
			operation.Dollar{
				Superunit: 3,
				Subunit:   5,
			},
		},
		// {
		// 	"ERROR: different results",
		// 	input{
		// 		operation.Dollar{
		// 			Superunit: 1,
		// 		},
		// 		operation.Dollar{
		// 			Superunit: 1,
		// 		},
		// 	},
		// 	operation.Dollar{
		// 		Superunit: 3,
		// 	},
		// },
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if actual := operation.SumDollars(test.input.a, test.input.b); actual != test.expected {
				t.Fatalf("expected %#v, actual %#v", test.expected, actual)
			}
		})
	}
}

func TestConcatenateSlice_Equally(t *testing.T) {
	t.Parallel()

	type (
		input struct {
			a []int
			b []int
		}
	)

	tests := []struct {
		name     string
		input    input
		expected []int
	}{
		{
			"OK",
			input{
				[]int{1, 2},
				[]int{-1, -2},
			},
			[]int{1, 2, -1, -2},
		},
		// {
		// 	"ERROR: different length",
		// 	input{
		// 		[]int{1, 2},
		// 		[]int{-1, -2},
		// 	},
		// 	[]int{1, -1, -2},
		// },
		// {
		// 	"ERROR: different values",
		// 	input{
		// 		[]int{1, 2},
		// 		[]int{-1, -2},
		// 	},
		// 	[]int{1, 2, 3, 4},
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := operation.ConcatenateSlice(test.input.a, test.input.b)

			if len(actual) != len(test.expected) {
				t.Fatalf("result lenghts are different: expected %d, actual %d", len(test.expected), len(actual))
			}

			for i, v := range actual {
				if v != test.expected[i] {
					t.Fatalf("values at %d index are different: expected %d, actual %d", i, test.expected[i], v)
				}
			}
		})
	}
}

func TestConcatenateSlice_Semantically(t *testing.T) {
	t.Parallel()

	type (
		input struct {
			a []int
			b []int
		}
	)

	tests := []struct {
		name     string
		input    input
		expected []int
	}{
		{
			"OK",
			input{
				[]int{1, 2},
				[]int{-1, -2},
			},
			[]int{-1, -2, 1, 2},
		},
		// {
		// 	"ERROR: missing 0",
		// 	input{
		// 		[]int{1, 2},
		// 		[]int{-1, -2},
		// 	},
		// 	[]int{-1, 0, 1, 2},
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := operation.ConcatenateSlice(test.input.a, test.input.b)

			if len(actual) != len(test.expected) {
				t.Fatalf("result lenghts are different: expected %d, actual %d", len(test.expected), len(actual))
			}

			found := make(map[int]struct{})

			for _, v := range actual {
				found[v] = struct{}{}
			}

			for _, v := range test.expected {
				if _, ok := found[v]; !ok {
					t.Fatalf("value %d is missing in the result", v)
				}

				delete(found, v)
			}

			if len(found) != 0 {
				t.Fatal("result does not contain expected values")
			}
		})
	}
}

func TestConcatenateSlice_Reflect(t *testing.T) {
	t.Parallel()

	type (
		input struct {
			a []int
			b []int
		}
	)

	tests := []struct {
		name     string
		input    input
		expected []int
	}{
		{
			"OK",
			input{
				[]int{1, 2},
				[]int{-1, -2},
			},
			[]int{1, 2, -1, -2},
		},
		// {
		// 	"ERROR: not same order",
		// 	input{
		// 		[]int{1, 2},
		// 		[]int{-1, -2},
		// 	},
		// 	[]int{-1, -2, 1, 2},
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := operation.ConcatenateSlice(test.input.a, test.input.b)

			if !reflect.DeepEqual(test.expected, actual) {
				t.Fatalf("values are not the same: expected %#v, actual %#v", test.expected, actual)
			}
		})
	}
}

func TestCmp_ApproximatePi(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input float64
	}{
		{
			"OK",
			math.Pi,
		},
		{
			"OK",
			3.14159,
		},
		{
			"OK",
			3.1415,
		},
		{
			"OK",
			3.14159265359,
		},
		// {
		// 	"ERROR: diverged",
		// 	3.141,
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if !cmp.Equal(math.Pi, test.input, cmpopts.EquateApprox(0, 0.0001)) {
				t.Fatalf("value diverge too far: %f", test.input)
			}
		})
	}
}

func TestCmp_SlicesSematicallySorting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			"OK",
			[]int{1, 2, 3},
			[]int{2, 1, 3},
		},
		// {
		// 	"ERROR",
		// 	[]int{1, 2, 3},
		// 	[]int{0, 1, 3},
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			opt := cmpopts.SortSlices(func(a, b int) bool {
				return a < b
			})

			if !cmp.Equal(test.input, test.output, opt) {
				t.Fatalf("values are not the same %s", cmp.Diff(test.input, test.output, opt))
			}
		})
	}
}

func TestCmp_MapsSematicallySorting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  map[int]string
		output map[int]string
	}{
		{
			"OK",
			map[int]string{1: "1", 2: "2"},
			map[int]string{2: "2", 1: "1"},
		},
		// {
		// 	"ERROR",
		// 	map[int]string{1: "1", 2: "2"},
		// 	map[int]string{1: "1", 3: "2"},
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			opt := cmpopts.SortMaps(func(a, b int) bool {
				return a < b
			})

			if !cmp.Equal(test.input, test.output, opt) {
				t.Fatalf("values are not the same %s", cmp.Diff(test.input, test.output, opt))
			}
		})
	}
}

func TestMessage_Equal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  operation.Message
		output operation.Message
	}{
		{
			"OK",
			"HELLO WORLD",
			"hello world",
		},
		// {
		// 	"ERROR",
		// 	"hello world",
		// 	"hola mundo",
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if !cmp.Equal(test.input, test.output) {
				t.Fatalf("values are not the same %s", cmp.Diff(test.input, test.output))
			}
		})
	}
}

func TestAlert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  operation.Alert
		output operation.Alert
	}{
		{
			"OK",
			operation.Alert{
				Message: "HELLO WORLD",
			},
			operation.Alert{
				Message: "hello world",
			},
		},
		// {
		// 	"ERROR",
		// 	operation.Alert{
		// 		Message: "hello world",
		// 	},
		// 	operation.Alert{
		// 		Message: "hola mundo",
		// 	},
		// },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if !cmp.Equal(test.input, test.output, cmpopts.IgnoreUnexported(operation.Alert{})) {
				t.Fatalf("values are not the same %s", cmp.Diff(test.input, test.output, cmpopts.IgnoreUnexported(operation.Alert{})))
			}
		})
	}
}
