// Manager package implements interface and functions for Queue management
// Queue is thread safe and implemented as FIFO list and HashMaps
//
// FIFO list should be mandatory, to prevent skipping of tasks,
// if there is more requests then infrastructure can handle.
//
// Hash Maps are used for faster Task management where we know exact task Id.
//
// Disclaimer:
// In normal usage i would add multiple levels of logging (Error, Warning, Debug)
// with possibility to turn it off, to prevent log spamming.
package manager

import (
	"log"
	"sync"

	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"
)

type (
	// Manager interface implements functions for Queue manipulation.
	Manager interface {
		AddTask(typ string, input interface{}) (UUID, error)
		TakeTask() (Task, error)
		TakeTaskByType(types []string) (Task, error)
		FinishTask(id UUID, output interface{}) error
		RetryTask(id UUID) error
		GetTask(id UUID) (Task, error)
		Status() (wait, progress, done int)
		DoneList() []Task
	}

	// manager struct contains data structures for Queue manipulation.
	manager struct {
		wait     *Queue
		progress map[UUID]Task
		done     map[UUID]Task

		sync.Mutex
	}
)

func NewManager() Manager {
	queue := NewQueue()

	return &manager{
		wait:     queue,
		progress: make(map[UUID]Task, 0),
		done:     make(map[UUID]Task, 0),
	}
}

// AddTask function will add task to waiting queue.
func (m *manager) AddTask(typ string, input interface{}) (UUID, error) {
	if typ == "" {
		return "", ErrMissingType
	}

	if input == nil {
		return "", ErrMissingInput
	}

	// Generate new Task.
	t := NewTask()
	t.Type = typ
	t.Input = input

	// Add task to Wait queue.
	m.Lock()
	m.wait.Push(&t)
	m.Unlock()

	// Log request.
	log.Println("AddTask:", t)

	// Return task id.
	return t.Id, nil
}

// TakeTask function will return task from wait queue to client/worker and move it to progress queue.
func (m *manager) TakeTask() (Task, error) {
	m.Lock()
	defer m.Unlock()

	// Pop task from wait queue.
	t := m.wait.Pop()
	if t == nil {
		return Task{}, ErrMissingTask
	}

	// Move task to progress queue.
	m.progress[t.Id] = *t

	// Log request.
	log.Println("TakeTask:", t)

	// Return task.
	return *t, nil
}

// TakeTaskByType function will return task by type, from wait queue to client/worker and move it to progress queue.
func (m *manager) TakeTaskByType(types []string) (Task, error) {
	m.Lock()
	defer m.Unlock()

	// Pop task by type from wait queue.
	t := m.wait.PopByType(types)
	if t == nil {
		return Task{}, ErrMissingTask
	}

	// Move task to progress queue.
	m.progress[t.Id] = *t

	// Log request.
	log.Println("TakeTaskByType:", t)

	// Return task.
	return *t, nil
}

// FinishTask function will save output from client/worker to task in progress and move it to done queue.
func (m *manager) FinishTask(id UUID, output interface{}) error {
	m.Lock()
	defer m.Unlock()

	// Take task by id from progress queue.
	if t, ok := m.progress[id]; ok {

		// Add output result to task.
		t.Output = output

		// Add task to done queue.
		m.done[t.Id] = t

		// Delete task from progress queue.
		delete(m.progress, id)

		// Log request.
		log.Println("FinishTask:", t)

		return nil
	}

	return ErrMissingTask
}

// RetryTask function will move task back to wait queue.
func (m *manager) RetryTask(id UUID) error {
	m.Lock()
	defer m.Unlock()

	// Take task by id from progress queue.
	if t, ok := m.progress[id]; ok {

		// Add task to wait queue.
		m.wait.Push(&t)

		// Delete task from progress queue.
		delete(m.progress, id)

		// Log request.
		log.Println("RetryTask:", t)

		return nil
	}

	return ErrMissingTask
}

// GetTask function returns task from queue and removes it from queue manager.
func (m *manager) GetTask(id UUID) (Task, error) {
	m.Lock()
	defer m.Unlock()

	// Take task by id from done queue
	if t, ok := m.done[id]; ok {

		// Delete task from done queue
		delete(m.done, id)

		// Log request.
		log.Println("GetTask:", t)

		return t, nil
	}

	return Task{}, ErrMissingTask
}

// Status function returns count for all queues.
func (m *manager) Status() (wait, progress, done int) {
	m.Lock()
	defer m.Unlock()

	return m.wait.Len(), len(m.progress), len(m.done)
}

// DoneList return list of finished tasks.
func (m *manager) DoneList() []Task {
	m.Lock()
	defer m.Unlock()

	ret := make([]Task, 0)
	for _, v := range m.done {
		ret = append(ret, v)
	}
	return ret
}
