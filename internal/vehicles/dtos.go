package vehicles

import (
	"time"
)

type CreateVehicleInput struct {
	Title string
	Type  string
	KM    int
}

type EditVehicleInput struct {
	ID    int
	Title string
	Type  string
	KM    int
}

type AddLogInput struct {
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
}

type EditLogInput struct {
	ID        int
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
}
