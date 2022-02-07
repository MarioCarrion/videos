package sum

import "testing"

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		inputA int
		inputB int
		output int
	}{
		{
			"When both are same",
			1,
			1,
			0,
		},
		{
			"When values are different",
			1,
			2,
			3,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Add(test.inputA, test.inputB)
			if got != test.output {
				t.Errorf("got %d, want %d", got, test.expecjted)
			}
		})
	}
}
