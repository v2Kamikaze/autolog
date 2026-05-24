package vehicleusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type DeleteLogEntryInput struct {
	VehicleID      int
	LogEntryID int
}

type DeleteLogEntryResponse struct {
	VehicleID         int
	TotalCost     int
	LogEntryCount int
	LogEntry      *entities.LogEntry
}

type DeleteLogEntryUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewDeleteLogEntryUseCase(vehiclePersistence persistence.VehiclePersistence) *DeleteLogEntryUseCase {
	return &DeleteLogEntryUseCase{
		vehiclePersistence: vehiclePersistence,
	}
}

func (u *DeleteLogEntryUseCase) Execute(ctx context.Context, input DeleteLogEntryInput) (*DeleteLogEntryResponse, error) {
	logEntry, err := u.vehiclePersistence.DeleteLogEntry(ctx, input.VehicleID, input.LogEntryID)
	if err != nil {
		return nil, err
	}

	vehicle, err := u.vehiclePersistence.GetVehicleByID(ctx, input.VehicleID)
	if err != nil {
		return nil, err
	}

	return &DeleteLogEntryResponse{
		VehicleID:         vehicle.ID,
		TotalCost:     vehicle.TotalCost(),
		LogEntryCount: vehicle.LogEntryCount(),
		LogEntry:      logEntry,
	}, nil
}
