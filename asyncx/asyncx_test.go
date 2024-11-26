package asyncx

import "testing"

func TestAsync(t *testing.T) {
	WithGoroutine(func() {
		t.Log("print")
	}, 3)
}
