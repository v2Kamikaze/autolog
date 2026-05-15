package carusecases

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type AddMaintenanceInput struct {
	CarID int
	Name  string
	Date  time.Time
	KM    int
	Cost  int
}

type AddMaintenanceResponse struct {
	CarID       int
	Maintenance *entities.Maintenance
	TotalCost   int
}

type AddMaintenanceUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewAddMaintenanceUseCase(carPersistence persistence.CarPersistence) *AddMaintenanceUseCase {
	return &AddMaintenanceUseCase{carPersistence: carPersistence}
}

func (u *AddMaintenanceUseCase) Execute(ctx context.Context, input AddMaintenanceInput) (*AddMaintenanceResponse, error) {
	maintenance := &entities.Maintenance{
		Name: input.Name,
		Date: input.Date,
		KM:   input.KM,
		Cost: input.Cost,
	}

	inserted, err := u.carPersistence.AddMaintenance(ctx, input.CarID, maintenance)
	if err != nil {
		return nil, err
	}

	car, err := u.carPersistence.GetCarByID(ctx, input.CarID)
	if err != nil {
		return nil, err
	}

	return &AddMaintenanceResponse{
		CarID:       input.CarID,
		Maintenance: inserted,
		TotalCost:   car.TotalCost(),
	}, nil
}
