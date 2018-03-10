package worker

import (
	"time"
)

// Host configures worker for specific host
func Host(host string) Option {
	return func(w *worker) {
		w.host = host
	}
}

// Delay sets interval between RPC server collections.
func Delay(delay time.Duration) Option {
	return func(w *worker) {
		w.delay = delay
	}
}
