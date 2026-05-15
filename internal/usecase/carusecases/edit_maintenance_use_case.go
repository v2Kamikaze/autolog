package carusecases

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type EditMaintenanceInput struct {
	CarID         int
	MaintenanceID int
	Name          string
	Date          time.Time
	KM            int
	Cost          int
}

type EditMaintenanceResponse struct {
	CarID       int
	Maintenance *entities.Maintenance
	TotalCost   int
}

type EditMaintenanceUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewEditMaintenanceUseCase(carPersistence persistence.CarPersistence) *EditMaintenanceUseCase {
	return &EditMaintenanceUseCase{carPersistence: carPersistence}
}

func (u *EditMaintenanceUseCase) Execute(ctx context.Context, input EditMaintenanceInput) (*EditMaintenanceResponse, error) {
	patch := &entities.Maintenance{
		Name: input.Name,
		Date: input.Date,
		KM:   input.KM,
		Cost: input.Cost,
	}

	updated, err := u.carPersistence.EditMaintenance(ctx, input.CarID, input.MaintenanceID, patch)
	if err != nil {
		return nil, err
	}

	car, err := u.carPersistence.GetCarByID(ctx, input.CarID)
	if err != nil {
		return nil, err
	}

	return &EditMaintenanceResponse{
		CarID:       input.CarID,
		Maintenance: updated,
		TotalCost:   car.TotalCost(),
	}, nil
}
