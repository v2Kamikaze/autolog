package persistence

import "errors"

var (
	ErrCarNotFound         = errors.New("car not found")
	ErrNilCar              = errors.New("nil car")
	ErrNilMaintenance      = errors.New("nil maintenance")
	ErrMaintenanceNotFound = errors.New("maintenance not found")
)
