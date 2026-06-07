package vehicles

import "time"

type CreateVehicleParams struct {
	Brand        string
	Model        string
	Year         int
	Version      string
	Engine       string
	Transmission TransmissionType
	Fuel         FuelType
	KM           int
}

type EditVehicleParams struct {
	ID           int
	Brand        string
	Model        string
	Year         int
	Version      string
	Engine       string
	Transmission TransmissionType
	Fuel         FuelType
	KM           int
}

type AddLogParams struct {
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
	Notes     string
}

type EditLogParams struct {
	ID        int
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
	Notes     string
}
