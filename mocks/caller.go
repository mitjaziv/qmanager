package mocks

import (
	. "github.com/mitjaziv/qmanager/internal/structures"
)

type callerMock struct {
	out interface{}
	err error
}

func NewCallerMock() *callerMock {
	return &callerMock{}
}

func (c *callerMock) Call(t *Task) error {
	t.Output = c.out
	return c.err
}

func (c *callerMock) SetError(err error) {
	c.err = err
}

func (c *callerMock) SetOutput(out interface{}) {
	c.out = out
}
