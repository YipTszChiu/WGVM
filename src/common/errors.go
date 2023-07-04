package common

import "errors"

var (
	ErrUnexpectedEnd = errors.New("unexpected end of section or function")
	ErrIntTooLong    = errors.New("integer representation too long")
	ErrIntTooLarge   = errors.New("integer too large")
	//errLenOutOfBounds = errors.New("length out of bounds")
)
