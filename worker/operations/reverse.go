// Reverse text resolver: Takes mirrored text and returns it reversed.
package operations

import (
	. "github.com/mitjaziv/qmanager/internal/structures"
)

type ReverseText struct {
}

func (rt *ReverseText) Call(t *Task) error {
	if t == nil {
		return ErrMissingArgs
	}

	text := t.Input.(string)

	// Reverse text.
	pos := 0
	ret := make([]byte, len(text))
	for i := len(text); i > 0; i-- {
		pos += copy(ret[pos:], text[i-1:i])
	}

	// Write reverse text to output
	t.Output = string(ret)

	return nil
}
