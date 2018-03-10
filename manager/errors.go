package manager

import (
	"errors"
)

var (
	ErrMissingType  = errors.New("missing task type")
	ErrMissingInput = errors.New("missing input parameters")
	ErrMissingTask  = errors.New("missing task in queue")
)
