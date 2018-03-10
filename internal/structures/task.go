package structures

import (
	"fmt"

	"github.com/mitjaziv/qmanager/internal/uuid"
)

type (
	Task struct {
		Id     uuid.UUID
		Type   string
		Input  interface{}
		Output interface{}
	}
)

var (
	// TaskTypes map contains all available task
	TaskTypes = map[string]bool{
		"fibonacci":  true,
		"arithmetic": true,
		"reverse":    true,
		"encoder":    true,
	}
)

func NewTask() Task {
	return Task{
		Id: uuid.NewUUID(),
	}
}

// String function implementation for Stringer interface.
func (t Task) String() string {
	return fmt.Sprintf("ID: %v\tType: %v\tInput: %v\tOutput: %v", t.Id, t.Type, t.Input, t.Output)
}
