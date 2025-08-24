package example_test

import (
	"errors"
	"testing"
	"testing/synctest"

	example "github.com/MarioCarrion/videos/2025/go-1-25-0/04-synctest"
)

func Test_Retry(t *testing.T) {
	t.Parallel()

	// Without "synctest", this test would take 2 seconds to run because of the two
	// retries with a 1 second sleep each. Using synctest, the test runs instantly.
	synctest.Test(t, func(t *testing.T) { // Remove this function to see how it was before this feature.
		mock := &mockPusher{
			errs: []error{
				errors.New("first"),
				errors.New("second"),
				nil,
			},
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

func Test_Calculate(t *testing.T) {
	t.Parallel()

	// Four different possible scenarios:
	// 1. Current code: it works because of "Wait"
	// 2. Commenting out "Wait" will fail, won't detect race condition unless run with "-race"
	// 3. Not using "synctest" will fail because of not waiting the goroutine to finish
	// 4. Using "-race" detects the race condition

	synctest.Test(t, func(t *testing.T) {
		var got int64

		go func() {
			got = example.Calculate(2)
		}()

		synctest.Wait()

		if want := int64(84); got != want {
			t.Fatalf("got: %v, want: %v", got, want)
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
