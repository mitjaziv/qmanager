package manager

import (
	"errors"
	"testing"

	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"

	"github.com/mitjaziv/qmanager/mocks"
)

func Test_RpcHandler(t *testing.T) {
	// managerMock
	managerMock := mocks.NewManagerMock()

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}
}

func Test_RpcHandler_AddTask(t *testing.T) {
	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Input = "test_input"

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(taskMock.Id, nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call AddTask.
	err := h.AddTask(taskMock, &replay)
	if err != nil {
		t.Error("expected nil error")
	}

	// get id
	id := replay.(UUID)
	if id != taskMock.Id {
		t.Error("expected id's to match")
	}
}

func Test_RpcHandler_AddTask_Error(t *testing.T) {
	// errorMock
	errorMock := errors.New("test error")

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(NewUUID(), errorMock)

	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Input = "test_input"

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call AddTask.
	err := h.AddTask(taskMock, &replay)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_RpcHandler_GetTask(t *testing.T) {
	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Input = "test_input"

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(taskMock, nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call GetTask.
	err := h.GetTask(taskMock.Id, &replay)
	if err != nil {
		t.Error("expected nil error")
	}

	// get task
	task := replay.(Task)
	if taskMock != task {
		t.Error("expected tasks to match")
	}
}

func Test_RpcHandler_GetTask_Error(t *testing.T) {
	// errorMock
	errorMock := errors.New("test error")

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(Task{}, errorMock)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call GetTask.
	err := h.GetTask("test_id", &replay)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_RpcHandler_TakeTask(t *testing.T) {
	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Input = "test_input"

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(taskMock, nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call TakeTask
	err := h.TakeTask(struct{}{}, &replay)
	if err != nil {
		t.Error("expected nil error")
	}

	// get task
	task := replay.(Task)
	if taskMock != task {
		t.Error("expected tasks to match")
	}
}

func Test_RpcHandler_TakeTask_Error(t *testing.T) {
	// errorMock
	errorMock := errors.New("test error")

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(Task{}, errorMock)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call TakeTask.
	err := h.TakeTask(struct{}{}, &replay)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_RpcHandler_TakeTask_MissingTask(t *testing.T) {
	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(Task{}, ErrMissingTask)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call TakeTask.
	err := h.TakeTask(struct{}{}, &replay)
	if err == ErrMissingTask {
		t.Error("expected missing task error")
	}
}

func Test_RpcHandler_TakeTaskByType(t *testing.T) {
	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Input = "test_input"

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(taskMock, nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call TakeTaskByType
	err := h.TakeTaskByType([]string{"type"}, &replay)
	if err != nil {
		t.Error("expected nil error")
	}

	// get task
	task := replay.(Task)
	if taskMock != task {
		t.Error("expected tasks to match")
	}
}

func Test_RpcHandler_TakeTaskByType_Error(t *testing.T) {
	// errorMock
	errorMock := errors.New("test error")

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(Task{}, errorMock)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call TakeTaskByType
	err := h.TakeTaskByType([]string{}, &replay)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_RpcHandler_TakeTaskByType_MissingTask(t *testing.T) {
	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(Task{}, ErrMissingTask)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call TakeTaskByType
	err := h.TakeTaskByType([]string{}, &replay)
	if err == ErrMissingTask {
		t.Error("expected missing task error")
	}
}

func Test_RpcHandler_FinishTask(t *testing.T) {
	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Input = "test_input"

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call FinishTask
	err := h.FinishTask(taskMock, &replay)
	if err != nil {
		t.Error("expected nil error")
	}
	if replay != nil {
		t.Error("expected nil replay")
	}
}

func Test_RpcHandler_FinishTask_Error(t *testing.T) {
	// errorMock
	errorMock := errors.New("test error")

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(errorMock)

	// taskMock
	taskMock := NewTask()
	taskMock.Type = "test_type"
	taskMock.Output = "test_output"

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call FinishTask.
	err := h.FinishTask(taskMock, &replay)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_RpcHandler_RetryTask(t *testing.T) {
	// taskMock
	taskMock := NewTask()

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call RetryTask.
	err := h.RetryTask(taskMock.Id, &replay)
	if err != nil {
		t.Error("expected nil error")
	}
}

func Test_RpcHandler_RetryTask_Error(t *testing.T) {
	// errorMock
	errorMock := errors.New("test error")

	// taskMock
	taskMock := NewTask()

	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(errorMock)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// reply
	var replay interface{}

	// call RetryTask.
	err := h.RetryTask(taskMock.Id, &replay)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_RpcHandler_RegisterHandler(t *testing.T) {
	// managerMock
	managerMock := mocks.NewManagerMock()
	managerMock.Return(nil)

	// handler
	h := NewRpcHandler(managerMock)
	if h == nil {
		t.Error("expected not nil handler")
	}

	// call FinishTask
	err := h.RegisterHandler()
	if err != nil {
		t.Error("expected nil error")
	}
}
