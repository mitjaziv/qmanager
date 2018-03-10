package operations

import (
	"errors"
)

var (
	ErrUnSupportedOp = errors.New("unsupported operation")
	ErrMissingArgs   = errors.New("missing arguments")
	ErrUnknownType   = errors.New("unknown type")
)
