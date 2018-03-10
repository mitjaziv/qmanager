// Text encoder: Takes string of text and returns BCrypt encrypted hash.
package operations

import (
	"golang.org/x/crypto/bcrypt"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

type (
	BCrypt struct {
	}

	bcInput struct {
		text string
		cost int
	}
)

// Call performs operation.
func (b *BCrypt) Call(t *Task) error {
	if t == nil {
		return ErrMissingArgs
	}
	i := t.Input.(string)

	// Generate hash
	hash, err := bcrypt.GenerateFromPassword([]byte(i), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Write hash to output variable.
	t.Output = string(hash)

	return nil
}
