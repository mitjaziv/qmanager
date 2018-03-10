package mocks

import (
	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"
)

// mock
type managerMock struct {
	mocks []interface{}
}

func NewManagerMock() *managerMock {
	return &managerMock{}
}

func (m *managerMock) AddTask(typ string, input interface{}) (UUID, error) {
	args := m.Called()
	if args[1] == nil {
		return args[0].(UUID), nil
	}
	return args[0].(UUID), args[1].(error)
}

func (m *managerMock) TakeTask() (Task, error) {
	args := m.Called()
	if args[1] == nil {
		return args[0].(Task), nil
	}
	return args[0].(Task), args[1].(error)
}

func (m *managerMock) TakeTaskByType(types []string) (Task, error) {
	args := m.Called()
	if args[1] == nil {
		return args[0].(Task), nil
	}
	return args[0].(Task), args[1].(error)
}

func (m *managerMock) FinishTask(id UUID, output interface{}) error {
	args := m.Called()
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *managerMock) RetryTask(id UUID) error {
	args := m.Called()
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *managerMock) GetTask(id UUID) (Task, error) {
	args := m.Called()
	if args[1] == nil {
		return args[0].(Task), nil
	}
	return args[0].(Task), args[1].(error)
}

func (m *managerMock) Status() (wait, progress, done int) {
	args := m.Called()
	return args[0].(int), args[1].(int), args[2].(int)
}

func (m *managerMock) DoneList() []Task {
	args := m.Called()
	return args[0].([]Task)
}

func (m *managerMock) Called() []interface{} {
	return m.mocks
}

func (m *managerMock) Return(values ...interface{}) {
	m.mocks = values
}
