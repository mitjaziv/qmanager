package uuid

import (
	"testing"
)

var (
	uuid14 = "4"
	uuid19 = []string{"8", "9", "a", "b"}
)

func Test_UUID(t *testing.T) {
	// Generate UUID - 1 and check for error and length.
	id1 := NewUUID()
	if len(id1) != 36 {
		t.Error("expected UUID length to be 36")
	}

	// UUID Version 4, needs to have {4} on position 14 and {8, 9, a, b} on position 19
	char14 := string(id1[14])
	char19 := string(id1[19])
	if char14 != uuid14 {
		t.Error("expected 4 on position 14")
	}
	if !(char19 == uuid19[0] ||
		char19 == uuid19[1] ||
		char19 == uuid19[2] ||
		char19 == uuid19[3]) {

		t.Error("expected {8, 9, a or b} on position 19")
	}

	// Generate UUID - 2 and check for error and length.
	id2 := NewUUID()
	if len(id2) != 36 {
		t.Error("expected UUID length to be 36")
	}

	// UUID Version 4, needs to have {4} on position 14 and {8, 9, a, b} on position 19
	char14 = string(id2[14])
	char19 = string(id2[19])
	if char14 != uuid14 {
		t.Error("expected 4 on position 14")
	}
	if !(char19 == uuid19[0] ||
		char19 == uuid19[1] ||
		char19 == uuid19[2] ||
		char19 == uuid19[3]) {

		t.Error("expected {8, 9, a or b} on position 19")
	}

	// Check for UUID difference.
	if id1 == id2 {
		t.Error("expected UUID's not to be equal")
	}
}
