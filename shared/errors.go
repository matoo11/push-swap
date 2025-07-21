package shared

import "errors"

// TODO: Define custom error types
var (
	ErrInvalidNumber = errors.New("invalid number")
	ErrDuplicate     = errors.New("duplicate number")
	ErrInvalidOp     = errors.New("invalid operation")
)
