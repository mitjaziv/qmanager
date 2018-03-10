package structures

import (
	"fmt"

	"testing"
)

func Test_NewTask(t *testing.T) {
	// NewTask
	task := NewTask()
	if task == (Task{}) {
		t.Error("expected Task not to be empty")
	}
	if task.Id == "" {
		t.Error("expected Task ID not to be empty")
	}

	// Output
	fmt.Println(task)
}
