package vehicles

import (
	"time"
)

type CreateVehicleParams struct {
	Title string
	Type  string
	KM    int
}

type EditVehicleParams struct {
	ID    int
	Title string
	Type  string
	KM    int
}

type AddLogParams struct {
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
}

type EditLogParams struct {
	ID        int
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
}
