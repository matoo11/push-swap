package shared

import "errors"

var (
	ErrInvalidNumber = errors.New("invalid number")
	ErrDuplicate     = errors.New("duplicate number")
	ErrInvalidOp     = errors.New("invalid operation")
)
