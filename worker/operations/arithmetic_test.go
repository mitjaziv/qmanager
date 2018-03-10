package operations

import (
	"testing"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

func Test_Arithmetic(t *testing.T) {
	// Test cases
	tests := []struct {
		input  string
		output float64
	}{
		{"( ( 55 + 5 ) + ( 2 * 5 ) )", 70},
		{"5 + 12 / 6 - 7 * 9", -56},
		{"( 17 + ( ( 8 - 2 ) * ( 5 * 6 ) ) )", 197},
		{"(17 + (8-2) * (5*6))", 197},
	}

	// ArithmeticSolver
	a := ArithmeticSolver{}

	// Run tests
	for _, row := range tests {
		task := Task{
			Input: row.input,
		}
		err := a.Call(&task)
		if err != nil {
			t.Errorf("expected nil error")
		}
		if task.Output != row.output {
			t.Errorf("expected result %f", row.output)
		}
	}
}

func Test_Arithmetic_Error(t *testing.T) {
	a := ArithmeticSolver{}

	// Task without input
	err := a.Call(nil)
	if err != ErrMissingArgs {
		t.Errorf("expected missing arguments")
	}

	// Task with incorrect input
	task := NewTask()
	task.Input = "5.ab + 12"

	err = a.Call(&task)
	if err == nil {
		t.Errorf("expected float parsing error")
	}

	// Unsupported operation
	_, err = calc("%", 6, 5)
	if err != ErrUnSupportedOp {
		t.Errorf("expected unsupported operation error")
	}
}
