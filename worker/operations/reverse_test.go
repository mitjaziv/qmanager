package operations

import (
	"testing"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

func Test_Reverse(t *testing.T) {
	// Test cases
	tests := []struct {
		input  interface{}
		output string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"ta suhi skafec pusca", "acsup cefaks ihus at"},
		{"2018", "8102"},
	}

	// ReverseText
	o := ReverseText{}

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
			t.Errorf("expected result %s", row.output)
		}
	}
}

func Test_Reverse_Error(t *testing.T) {
	o := ReverseText{}

	// Call Reverse text without input.
	err := o.Call(nil)
	if err != ErrMissingArgs {
		t.Errorf("expected missing arguments")
	}
}
