package interpreter

import "errors"

var (
	ErrTrap              = errors.New("unreachable")
	ErrCallStackOverflow = errors.New("call stack exhausted")
	ErrTypeMismatch      = errors.New("indirect call type mismatch")
	ErrUndefinedElem     = errors.New("undefined element")
	ErrUninitializedElem = errors.New("uninitialized element")
	ErrMemOutOfBounds    = errors.New("out of bounds memory access")
	ErrImmutableGlobal   = errors.New("immutable global")
	ErrIntOverflow       = errors.New("integer overflow")
	ErrConvertToInt      = errors.New("invalid conversion to integer")
)
