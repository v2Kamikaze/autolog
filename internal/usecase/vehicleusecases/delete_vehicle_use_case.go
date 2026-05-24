package vehicleusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type DeleteVehicleInput struct {
	ID int
}

type DeleteVehicleResponse struct {
	Vehicle *entities.Vehicle
}

type DeleteVehicleUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewDeleteVehicleUseCase(vehiclePersistence persistence.VehiclePersistence) *DeleteVehicleUseCase {
	return &DeleteVehicleUseCase{
		vehiclePersistence: vehiclePersistence,
	}
}

func (u *DeleteVehicleUseCase) Execute(ctx context.Context, input DeleteVehicleInput) (*DeleteVehicleResponse, error) {
	vehicle, err := u.vehiclePersistence.DeleteVehicle(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteVehicleResponse{Vehicle: vehicle}, nil
}
