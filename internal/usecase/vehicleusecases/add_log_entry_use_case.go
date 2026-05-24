package vehicleusecases

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type AddLogEntryInput struct {
	VehicleID int
	Name  string
	Type  entities.LogEntryCategory
	Date  time.Time
	KM    int
	Cost  int
}

type AddLogEntryResponse struct {
	VehicleID         int
	LogEntry      *entities.LogEntry
	TotalCost     int
	LogEntryCount int
}

type AddLogEntryUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewAddLogEntryUseCase(vehiclePersistence persistence.VehiclePersistence) *AddLogEntryUseCase {
	return &AddLogEntryUseCase{vehiclePersistence: vehiclePersistence}
}

func (u *AddLogEntryUseCase) Execute(ctx context.Context, input AddLogEntryInput) (*AddLogEntryResponse, error) {
	if !input.Type.Valid() {
		return nil, persistence.ErrInvalidLogEntryType
	}

	logEntry := &entities.LogEntry{
		Name: input.Name,
		Type: input.Type,
		Date: input.Date,
		KM:   input.KM,
		Cost: input.Cost,
	}

	inserted, err := u.vehiclePersistence.AddLogEntry(ctx, input.VehicleID, logEntry)
	if err != nil {
		return nil, err
	}

	vehicle, err := u.vehiclePersistence.GetVehicleByID(ctx, input.VehicleID)
	if err != nil {
		return nil, err
	}

	return &AddLogEntryResponse{
		VehicleID:         input.VehicleID,
		LogEntry:      inserted,
		TotalCost:     vehicle.TotalCost(),
		LogEntryCount: vehicle.LogEntryCount(),
	}, nil
}
