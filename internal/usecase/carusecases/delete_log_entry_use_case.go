package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type DeleteLogEntryInput struct {
	CarID      int
	LogEntryID int
}

type DeleteLogEntryResponse struct {
	CarID         int
	TotalCost     int
	LogEntryCount int
	LogEntry      *entities.LogEntry
}

type DeleteLogEntryUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewDeleteLogEntryUseCase(carPersistence persistence.CarPersistence) *DeleteLogEntryUseCase {
	return &DeleteLogEntryUseCase{
		carPersistence: carPersistence,
	}
}

func (u *DeleteLogEntryUseCase) Execute(ctx context.Context, input DeleteLogEntryInput) (*DeleteLogEntryResponse, error) {
	logEntry, err := u.carPersistence.DeleteLogEntry(ctx, input.CarID, input.LogEntryID)
	if err != nil {
		return nil, err
	}

	car, err := u.carPersistence.GetCarByID(ctx, input.CarID)
	if err != nil {
		return nil, err
	}

	return &DeleteLogEntryResponse{
		CarID:         car.ID,
		TotalCost:     car.TotalCost(),
		LogEntryCount: car.LogEntryCount(),
		LogEntry:      logEntry,
	}, nil
}
