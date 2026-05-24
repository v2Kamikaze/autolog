package vehicleusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type EditVehicleInput struct {
	ID    int
	Title string
	Type  string
	KM    int
}

type EditVehicleResponse struct {
	Vehicle *entities.Vehicle
}

type EditVehicleUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewEditVehicleUseCase(vehiclePersistence persistence.VehiclePersistence) *EditVehicleUseCase {
	return &EditVehicleUseCase{
		vehiclePersistence: vehiclePersistence,
	}
}

func (uc *EditVehicleUseCase) Execute(ctx context.Context, input EditVehicleInput) (*EditVehicleResponse, error) {
	patch := &entities.Vehicle{
		ID:    input.ID,
		Title: input.Title,
		Type:  input.Type,
		KM:    input.KM,
	}
	vehicle, err := uc.vehiclePersistence.EditVehicle(ctx, patch)
	if err != nil {
		return nil, err
	}
	return &EditVehicleResponse{Vehicle: vehicle}, nil
}
