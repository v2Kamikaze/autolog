package vehicleusecases

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type EditLogEntryInput struct {
	VehicleID      int
	LogEntryID int
	Name       string
	Type       entities.LogEntryCategory
	Date       time.Time
	KM         int
	Cost       int
}

type EditLogEntryResponse struct {
	VehicleID    int
	LogEntry *entities.LogEntry
	TotalCost int
}

type EditLogEntryUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewEditLogEntryUseCase(vehiclePersistence persistence.VehiclePersistence) *EditLogEntryUseCase {
	return &EditLogEntryUseCase{vehiclePersistence: vehiclePersistence}
}

func (u *EditLogEntryUseCase) Execute(ctx context.Context, input EditLogEntryInput) (*EditLogEntryResponse, error) {
	if !input.Type.Valid() {
		return nil, persistence.ErrInvalidLogEntryType
	}

	patch := &entities.LogEntry{
		Name: input.Name,
		Type: input.Type,
		Date: input.Date,
		KM:   input.KM,
		Cost: input.Cost,
	}
	updated, err := u.vehiclePersistence.EditLogEntry(ctx, input.VehicleID, input.LogEntryID, patch)
	if err != nil {
		return nil, err
	}

	vehicle, err := u.vehiclePersistence.GetVehicleByID(ctx, input.VehicleID)
	if err != nil {
		return nil, err
	}

	return &EditLogEntryResponse{
		VehicleID:     input.VehicleID,
		LogEntry:  updated,
		TotalCost: vehicle.TotalCost(),
	}, nil
}
