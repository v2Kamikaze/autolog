package mocks

import "errors"

var (
	ErrNilVehicle          = errors.New("nil vehicle")
	ErrVehicleNotFound     = errors.New("vehicle not found")
	ErrNilLogEntry         = errors.New("nil log entry")
	ErrLogEntryNotFound    = errors.New("log entry not found")
	ErrInvalidLogEntryType = errors.New("invalid log entry type")
)
