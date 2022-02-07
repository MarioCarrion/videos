package sum

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type input struct {
		a Value
		b Value
	}

	tests := []struct {
		name   string
		input  input
		output Value
	}{
		{
			"Both are the same",
			input{
				Value{A: 1, B: 2},
				Value{A: 1, B: 2},
			},
			Value{},
		},
		{
			"Values are different",
			input{
				Value{A: 1, B: 2},
				Value{A: 2, B: 2},
			},
			Value{A: 3, B: 4},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.input.a, tt.input.b)
			if got != tt.output {
				t.Fatalf("got %v, wanted %v", got, tt.output)
			}
		})
	}
}
