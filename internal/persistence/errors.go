package persistence

import "errors"

var (
	ErrNilCar               = errors.New("nil car")
	ErrCarNotFound          = errors.New("car not found")
	ErrNilLogEntry          = errors.New("nil log entry")
	ErrLogEntryNotFound     = errors.New("log entry not found")
	ErrInvalidLogEntryType  = errors.New("invalid log entry type")
)
