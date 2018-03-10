package manager

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HTTP(t *testing.T) {
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

	// HttpHandler
	h := NewHttpHandler(m)
	if h == nil {
		t.Error("expected handler not be nil")
	}

	// Register Handlers
	h.RegisterHandler()

	// Create recorder
	rec := httptest.NewRecorder()

	// Server HTTP
	h.ServeHTTP(rec, nil)

	// Check status
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
