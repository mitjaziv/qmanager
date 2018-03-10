package operations

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

func Test_BCrypt(t *testing.T) {
	o := BCrypt{}

	// task
	task := Task{}
	task.Input = "test_input"

	// Call BCrypt call
	err := o.Call(&task)
	if err != nil {
		t.Errorf("expected nil error")
	}

	// Check if hash matches password
	err = bcrypt.CompareHashAndPassword(
		[]byte(task.Output.(string)),
		[]byte(task.Input.(string)),
	)
	if err != nil {
		t.Errorf("expected nil error")
	}
}

func Test_BCrypt_Error(t *testing.T) {
	o := BCrypt{}

	// Call BCrypt without input.
	err := o.Call(nil)
	if err != ErrMissingArgs {
		t.Errorf("expected missing arguments")
	}
}

func Test_BCrypt_EmptyString(t *testing.T) {
	o := BCrypt{}

	// Test input
	input := bcInput{
		text: "test_input",
		cost: -1,
	}
	_ = input
	_ = o

	// Call BCrypt without input.
	// out, err := o.Call(input)

	/// fmt.Println(out, err)

	// assert.Error(t, err, "expected error")
	// assert.Nil(t, out, "expected nil for output")
}
