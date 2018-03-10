package worker

import (
	"testing"

	. "github.com/mitjaziv/qmanager/internal/structures"

	"github.com/mitjaziv/qmanager/mocks"
)

func Test_CallbackFactory(t *testing.T) {
	// mocks
	callerMock := mocks.NewCallerMock()
	callerMock.SetError(nil)
	callerMock.SetOutput("test_output")

	// new callback factory
	cb := NewCallbackFactory()
	if cb == nil {
		t.Error("expected new callback factory")
	}

	// register callback
	cb.Register(callerMock, "test")
	if cb == nil {
		t.Error("expected new callback factory")
	}

	// task
	task := Task{}
	task.Type = "test"

	// calls
	err := cb.Call(&task)
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Output != "test_output" {
		t.Error("expected different task output")
	}

	// task with unknown type
	task = Task{}
	task.Type = "unknown"

	// call
	err = cb.Call(&task)
	if err == nil {
		t.Error("expected nil error")
	}

	// registered types
	types := cb.Types()
	if len(types) != 1 && types[0] != "test" {
		t.Error("expected types to contain only 'test'")
	}

	// register already registered callback
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("register should have panicked")
			}
		}()
		// register should cause a panic
		cb.Register(callerMock, "test")
	}()
}
