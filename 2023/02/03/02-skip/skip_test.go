package skip_test

import (
	"testing"

	skip "github.com/MarioCarrion/videos/2023/02/03/02-skip"
)

func Test_String(t *testing.T) {
	t.Parallel()

	if got := skip.String(); got != "hello world" {
		t.Fatalf("got %v expected %v", got, "hello world")
	}
}

func Test_Integer(t *testing.T) {
	t.Parallel()

	if i := skip.Integer(); i != 42 {
		t.Fatalf("got %v expected %v", i, 42)
	}
}
