package carusecases

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type EditLogEntryInput struct {
	CarID      int
	LogEntryID int
	Name       string
	Type       entities.LogEntryCategory
	Date       time.Time
	KM         int
	Cost       int
}

type EditLogEntryResponse struct {
	CarID    int
	LogEntry *entities.LogEntry
	TotalCost int
}

type EditLogEntryUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewEditLogEntryUseCase(carPersistence persistence.CarPersistence) *EditLogEntryUseCase {
	return &EditLogEntryUseCase{carPersistence: carPersistence}
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
	updated, err := u.carPersistence.EditLogEntry(ctx, input.CarID, input.LogEntryID, patch)
	if err != nil {
		return nil, err
	}

	car, err := u.carPersistence.GetCarByID(ctx, input.CarID)
	if err != nil {
		return nil, err
	}

	return &EditLogEntryResponse{
		CarID:     input.CarID,
		LogEntry:  updated,
		TotalCost: car.TotalCost(),
	}, nil
}
