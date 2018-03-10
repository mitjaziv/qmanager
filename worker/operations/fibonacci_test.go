package operations

import (
	"testing"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

func Test_Fibonacci(t *testing.T) {
	// Test cases
	tests := []struct {
		input  interface{}
		output int
	}{
		{0, 0},
		{"1", 1},
		{2, 1},
		{"3", 2},
		{10, 55},
		{"20", 6765},
		{20, 6765},
	}

	// Fibonacci
	o := Fibonacci{}

	// Run tests
	for _, row := range tests {
		task := Task{
			Input: row.input,
		}
		err := o.Call(&task)
		if err != nil {
			t.Errorf("expected nil error")
		}
		if task.Output != row.output {
			t.Errorf("expected fibonacci result %d", row.output)
		}
	}
}

func Test_Fibonacci_Error(t *testing.T) {
	o := Fibonacci{}

	// Call Fibonacci without input.
	err := o.Call(nil)
	if err != ErrMissingArgs {
		t.Errorf("expected missing arguments")
	}
}
