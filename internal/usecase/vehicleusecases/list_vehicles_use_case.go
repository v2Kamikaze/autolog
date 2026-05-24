package vehicleusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type ListVehiclesResponse struct {
	Vehicles []*entities.Vehicle
}

type ListVehiclesUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewListVehiclesUseCase(vehiclePersistence persistence.VehiclePersistence) *ListVehiclesUseCase {
	return &ListVehiclesUseCase{vehiclePersistence: vehiclePersistence}
}

func (u *ListVehiclesUseCase) Execute(ctx context.Context) (*ListVehiclesResponse, error) {
	vehicles, err := u.vehiclePersistence.ListVehicles(ctx)
	if err != nil {
		return nil, err
	}
	return &ListVehiclesResponse{Vehicles: vehicles}, nil
}
