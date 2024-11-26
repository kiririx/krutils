package asyncx

import "testing"

func TestAsync(t *testing.T) {
	WithGoroutine(func() {
		t.Log("print")
	}, 3)
}

func TestTask(t *testing.T) {
	err := ScheduleTask("21 * * * 2", func() {
		t.Log("print")
	})
	if err != nil {
		return
	}
}
