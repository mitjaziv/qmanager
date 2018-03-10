// Fibonacci resolver: Takes an integer and returns the result of Fibonacci function.
package operations

import (
	"strconv"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

type Fibonacci struct {
}

// Fibonacci function
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Call performs operation.
func (rt *Fibonacci) Call(t *Task) error {
	if t == nil {
		return ErrMissingArgs
	}

	var num int
	var err error

	switch t.Input.(type) {
	case string:
		// Convert string to int
		num, err = strconv.Atoi(t.Input.(string))
		if err != nil {
			return err
		}
	case int:
		num = t.Input.(int)
	default:
		return ErrUnknownType

	}

	// Write fibonacci number to output
	t.Output = fibonacci(num)

	return nil
}
