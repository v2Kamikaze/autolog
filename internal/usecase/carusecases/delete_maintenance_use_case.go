package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type DeleteMaintenanceInput struct {
	CarID         int
	MaintenanceID int
}

type DeleteMaintenanceResponse struct {
	CarID       int
	TotalCost   int
	Maintenance *entities.Maintenance
}

type DeleteMaintenanceUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewDeleteMaintenanceUseCase(carPersistence persistence.CarPersistence) *DeleteMaintenanceUseCase {
	return &DeleteMaintenanceUseCase{
		carPersistence: carPersistence,
	}
}

func (u *DeleteMaintenanceUseCase) Execute(ctx context.Context, input DeleteMaintenanceInput) (*DeleteMaintenanceResponse, error) {
	maintenance, err := u.carPersistence.DeleteMaintenance(ctx, input.CarID, input.MaintenanceID)
	if err != nil {
		return nil, err
	}

	car, err := u.carPersistence.GetCarByID(ctx, input.CarID)
	if err != nil {
		return nil, err
	}

	return &DeleteMaintenanceResponse{
		CarID:       car.ID,
		TotalCost:   car.TotalCost(),
		Maintenance: maintenance,
	}, nil
}
