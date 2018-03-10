package manager

import (
	"testing"
)

func Test_Manager_NewManager(t *testing.T) {
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}
}

func Test_Manager_AddTask(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// AddTask with missing type
	_, err := m.AddTask("", nil)
	if err != ErrMissingType {
		t.Error("expected missing type error")
	}

	// AddTask with missing input
	_, err = m.AddTask("test", nil)
	if err != ErrMissingInput {
		t.Error("expected missing input error")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}
}

func Test_Manager_TakeTask(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// TakeTask when queue empty
	_, err := m.TakeTask()
	if err != ErrMissingTask {
		t.Error("expected missing task error")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}

	// TakeTask when queue not empty
	task, err := m.TakeTask()
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Id != id {
		t.Error("expcted tasks id to be equal")
	}
}

func Test_Manager_TakeTaskByType(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// TakeTask when queue empty
	_, err := m.TakeTaskByType([]string{"test_type"})
	if err != ErrMissingTask {
		t.Error("expected missing task error")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}

	// TakeTask when queue not empty
	task, err := m.TakeTaskByType([]string{"test_type"})
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Id != id {
		t.Error("expcted tasks id to be equal")
	}
}

func Test_Manager_FinishTask(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// FinishTask with empty queues
	err := m.FinishTask("test_id", "test_output")
	if err != ErrMissingTask {
		t.Error("expected missing task error")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}

	// TakeTask successful
	task, err := m.TakeTask()
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Id != id {
		t.Error("expcted tasks id to be equal")
	}

	// FinishTask successful
	err = m.FinishTask(task.Id, "test_output")
	if err != nil {
		t.Error("expected nil error")
	}
}

func Test_Manager_RetryTask(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// RetryTask with empty queues
	err := m.RetryTask("test_id")
	if err != ErrMissingTask {
		t.Error("expected missing task error")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}

	// TakeTask successful
	task, err := m.TakeTask()
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Id != id {
		t.Error("expcted tasks id to be equal")
	}

	// RetryTask successful
	err = m.RetryTask(task.Id)
	if err != nil {
		t.Error("expected nil error")
	}
}

func Test_Manager_GetTask(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// GetTask with empty queues
	_, err := m.GetTask("test_id")
	if err != ErrMissingTask {
		t.Error("expected missing task error")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}

	// TakeTask successful
	task, err := m.TakeTask()
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Id != id {
		t.Error("expcted tasks id to be equal")
	}

	// FinishTask
	err = m.FinishTask(task.Id, "test_output")
	if err != nil {
		t.Error("expected nil error")
	}

	// DoneList successful
	list := m.DoneList()
	if list == nil {
		t.Error("expected list not to be nil")
	}
	if list[0].Id != id {
		t.Error("expected task to be finished")
	}

	// GetTask successful
	task, err = m.GetTask(id)
	if err != nil {
		t.Error("expected nil error")
	}
	if task.Id == id && task.Output != "test_output" {
		t.Error("expected finished task")
	}
}

func Test_Manager_Status(t *testing.T) {
	// NewManager
	m := NewManager()
	if m == nil {
		t.Error("expected manager not be nil")
	}

	// AddTask successful
	id, err := m.AddTask("test_type", "test_input")
	if err != nil {
		t.Error("expected nil error")
	}
	if id == "" {
		t.Error("expected new task id")
	}

	w, p, d := m.Status()
	if w != 1 && p == 0 && d == 0 {
		t.Error("expected different statistics")
	}
}
