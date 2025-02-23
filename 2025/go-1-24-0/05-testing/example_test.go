package example_test

import (
	"testing"

	example "github.com/MarioCarrion/videos/2025/go-1-24-0/05-testing"
)

func TestRun(t *testing.T) {
	srv := example.NewServer()

	go func() {
		<-t.Context().Done()
		t.Log("context canceled")
		srv.Stop()
	}()

	t.Cleanup(func() {
		<-srv.Done()
		t.Log("server stopped")

		if !srv.Stopped() {
			t.Fatal("server not stopped")
		}
	})

	got := srv.Multiply(10, 2)
	if got != 20 {
		t.Fatalf("got: %d, want: 20", got)
	}

	t.Log("test finished")
}
