package carusecases

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type AddLogEntryInput struct {
	CarID int
	Name  string
	Type  entities.LogEntryCategory
	Date  time.Time
	KM    int
	Cost  int
}

type AddLogEntryResponse struct {
	CarID         int
	LogEntry      *entities.LogEntry
	TotalCost     int
	LogEntryCount int
}

type AddLogEntryUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewAddLogEntryUseCase(carPersistence persistence.CarPersistence) *AddLogEntryUseCase {
	return &AddLogEntryUseCase{carPersistence: carPersistence}
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

	inserted, err := u.carPersistence.AddLogEntry(ctx, input.CarID, logEntry)
	if err != nil {
		return nil, err
	}

	car, err := u.carPersistence.GetCarByID(ctx, input.CarID)
	if err != nil {
		return nil, err
	}

	return &AddLogEntryResponse{
		CarID:         input.CarID,
		LogEntry:      inserted,
		TotalCost:     car.TotalCost(),
		LogEntryCount: car.LogEntryCount(),
	}, nil
}
