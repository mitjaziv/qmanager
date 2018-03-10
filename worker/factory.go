package worker

import (
	"errors"
	"fmt"
	"log"

	. "github.com/mitjaziv/qmanager/internal/structures"
)

type (
	Caller interface {
		Call(t *Task) error
	}

	CallbackFactory struct {
		callbacks map[string]Caller
	}
)

func NewCallbackFactory() *CallbackFactory {
	return &CallbackFactory{
		callbacks: make(map[string]Caller),
	}
}

func (cf *CallbackFactory) Register(callback Caller, typ string) *CallbackFactory {
	if _, ok := cf.callbacks[typ]; ok {
		panic(fmt.Sprintf("registering multiple callbacks for 'type' %s.", typ))
	}
	cf.callbacks[typ] = callback
	log.Println("registered callback for:", typ)

	return cf
}

func (cf *CallbackFactory) Call(t *Task) error {
	h, ok := cf.callbacks[t.Type]
	if !ok {
		return errors.New("callback not found")
	}
	return h.Call(t)
}

func (cf *CallbackFactory) Types() []string {
	types := make([]string, 0, len(cf.callbacks))
	for key := range cf.callbacks {
		types = append(types, key)
	}
	return types
}
