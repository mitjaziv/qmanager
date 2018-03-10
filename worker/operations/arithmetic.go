// Basic arithmetic resolver: Takes basic arithmetic problem and returns the result.
//
// Disclaimer:
// In ordinary case I would implement this with, stack (LIFO) data type, not with slice hack.
// Using google API for evaluation of complex arithmetic would be even better.
package operations

import (
	"strconv"
	"strings"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

type (
	ArithmeticSolver struct {
	}
)

var (
	// Precedence map
	precedence = map[string]int{
		"(": 0,
		")": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
)

// process function takes input and process operations and values.
func process(input string) (out float64, err error) {
	stackO := make([]string, 0)
	stackV := make([]float64, 0)

	// Create tokens from input string.
	tokens := strings.Fields(input)

	// Temp values
	var op string
	var v1, v2 float64

	// Loop over tokens and evaluate arithmetic.
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// If token is value, parse it and push on stack.
		if _, ok := precedence[t]; !ok {
			v, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return 0, err
			}

			// Push value on stack.
			stackV = append([]float64{v}, stackV...)

			// Proceed to next token.
			continue
		}

		for {
			if len(stackO) == 0 ||
				t == "(" ||
				(precedence[t] > precedence[stackO[len(stackO)-1]]) {

				// Push operation on stack.
				stackO = append([]string{t}, stackO...)
				break
			}

			// Pop operation from stack.
			op, stackO = stackO[0], stackO[1:]

			// Brake on left parenthesis.
			if op == "(" {
				break
			}

			// Pop val from stack.
			v2, stackV = stackV[0], stackV[1:]

			// Pop val from stack.
			v1, stackV = stackV[0], stackV[1:]

			// Calculate actual operation
			val, err := calc(op, v1, v2)
			if err != nil {
				return 0, err
			}

			// Push value on stack.
			stackV = append([]float64{val}, stackV...)
		}
	}

	// Finish evaluating data on stacks.
	for {
		if len(stackO) == 0 {
			break
		}

		// Pop operation from stack.
		op, stackO = stackO[0], stackO[1:]

		// Pop val from stack.
		v2, stackV = stackV[0], stackV[1:]

		// Pop val from stack.
		v1, stackV = stackV[0], stackV[1:]

		// Calculate actual operation.
		val, err := calc(op, v1, v2)
		if err != nil {
			return 0, err
		}

		// Push value on stack.
		stackV = append([]float64{val}, stackV...)
	}

	// Get result from stack and returns it.
	return stackV[0], nil
}

// calc function performs operation on given operands.
func calc(op string, v1, v2 float64) (float64, error) {
	switch op {
	case "+":
		v1 = v1 + v2
	case "-":
		v1 = v1 - v2
	case "*":
		v1 = v1 * v2
	case "/":
		v1 = v1 / v2
	default:
		return 0, ErrUnSupportedOp
	}
	return v1, nil
}

// format input string.
func format(input string) string {
	i := 0
	for {
		if i >= len(input)-1 {
			break
		}

		// Get char on position
		c := string(input[i])
		c1 := string(input[i+1])

		// Add whitespace to operand if needed
		if _, ok := precedence[c]; ok && c1 != " " {
			input = input[:i+1] + " " + input[i+1:]
		}

		// Add whitespace to number if needed
		if _, ok := precedence[c1]; ok && c != " " {
			input = input[:i+1] + " " + input[i+1:]
		}

		i++
	}

	return input
}

// Call performs operation.
func (as *ArithmeticSolver) Call(t *Task) error {
	if t == nil {
		return ErrMissingArgs
	}
	i := t.Input.(string)

	// Format input, add needed whitespaces.
	i = format(i)

	// Calculate
	out, err := process(i)
	if err != nil {
		return err
	}

	// Write result to output.
	t.Output = out

	return nil
}
