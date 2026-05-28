package vehicles

import (
	"time"

	"github.com/v2code/autolog/internal/persistence/entities"
)

type CreateVehicleInput struct {
	Title string
	Type  string
	KM    int
}

type EditVehicleInput struct {
	Title string
	Type  string
	KM    int
}

type AddLogInput struct {
	VehicleID int
	Name      string
	Type      entities.LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
}

type EditLogInput struct {
	Name string
	Type entities.LogEntryCategory
	Date time.Time
	KM   int
	Cost int
}
