package example_test

import (
	"errors"
	"testing"
	"testing/synctest"

	example "github.com/MarioCarrion/videos/2025/go-1-24-0/05-testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	synctest.Run(func() { // Remove this function to see how it was before this experimental feature.
		mock := &mockPusher{
			errs: []error{errors.New("first"), errors.New("second"), nil},
		}

		gotErr := example.Retry(3, 10, mock)
		if gotErr != nil {
			t.Fatalf("got: %v, want: nil", gotErr)
		}
		if mock.retries != 3 {
			t.Fatalf("got: %d, want: 3", mock.retries)
		}
		if mock.value != 10 {
			t.Fatalf("got: %d, want: 10", mock.value)
		}
	})
}

type mockPusher struct {
	errs    []error
	retries int64
	value   int64
}

func (m *mockPusher) Push(value int64) error {
	m.value = value
	err := m.errs[m.retries]
	m.retries++
	return err
}
