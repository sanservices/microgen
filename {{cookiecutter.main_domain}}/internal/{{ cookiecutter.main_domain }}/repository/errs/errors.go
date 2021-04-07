package errs

import "errors"

var (
	// ErrNotAbleToStartTransaction is should be returned when transaction to db can't be started
	ErrNotAbleToStartTransaction = errors.New("can't start transaction to underlying datasource")
	// ErrNotImplemented is error for non-implemented methods
	ErrNotImplemented = errors.New("not implemented")
)
