package manager

import (
	"encoding/gob"
	"net/rpc"

	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"
)

type (
	RpcHandler struct {
		manager Manager
	}
)

func NewRpcHandler(m Manager) *RpcHandler {
	return &RpcHandler{
		manager: m,
	}
}

// AddTask function adds Task in argument to Queue manager.
func (h *RpcHandler) AddTask(t Task, reply *interface{}) error {
	id, err := h.manager.AddTask(t.Type, t.Input)
	if err != nil {
		return err
	}

	// Add id to reply interface
	*reply = id

	return nil
}

// GetTask function returns finished task with results and removes it from Queue manager.
func (h *RpcHandler) GetTask(id UUID, reply *interface{}) error {
	t, err := h.manager.GetTask(id)
	if err != nil {
		return err
	}

	// Add task to reply interface
	*reply = t

	return nil
}

// TakeTask function returns first Task in Queue manager and marks it as in progress.
func (h *RpcHandler) TakeTask(void struct{}, reply *interface{}) error {
	t, err := h.manager.TakeTask()
	if err == ErrMissingTask {
		return nil
	}
	if err != nil {
		return err
	}

	// Add task to reply interface
	*reply = t

	return nil
}

// TakeTask function returns first Task with matching type in Queue manager and marks it as in progress.
func (h *RpcHandler) TakeTaskByType(types []string, reply *interface{}) error {
	t, err := h.manager.TakeTaskByType(types)
	if err == ErrMissingTask {
		return nil
	}
	if err != nil {
		return err
	}

	// Add task to reply interface
	*reply = t

	return nil
}

// FinishTask function writes results in Task and marks it as done.
func (h *RpcHandler) FinishTask(t Task, reply *interface{}) error {
	err := h.manager.FinishTask(t.Id, t.Output)
	if err != nil {
		return err
	}
	return nil
}

// RetryTask function moves Task back to wait queue.
func (h *RpcHandler) RetryTask(id UUID, reply *interface{}) error {
	err := h.manager.RetryTask(id)
	if err != nil {
		return err
	}
	return nil
}

// RegisterHandler  registers RPC and HTTP handlers
func (h *RpcHandler) RegisterHandler() error {
	err := rpc.Register(h)
	if err != nil {
		return err
	}
	rpc.HandleHTTP()

	// Register types
	task := new(Task)
	uuid := new(UUID)

	gob.Register(*task)
	gob.Register(*uuid)

	return nil
}
