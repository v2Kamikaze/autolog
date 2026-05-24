package vehicleusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type CreateVehicleInput struct {
	Title string
	Type  string
	KM    int
}

type CreateVehicleResponse struct {
	Vehicle *entities.Vehicle
}

type CreateVehicleUseCase struct {
	vehiclePersistence persistence.VehiclePersistence
}

func NewCreateVehicleUseCase(vehiclePersistence persistence.VehiclePersistence) *CreateVehicleUseCase {
	return &CreateVehicleUseCase{vehiclePersistence: vehiclePersistence}
}

func (u *CreateVehicleUseCase) Execute(ctx context.Context, input CreateVehicleInput) (*CreateVehicleResponse, error) {
	vehicle := &entities.Vehicle{
		Title:        input.Title,
		Type:         input.Type,
		KM:           input.KM,
		LogEntries: []*entities.LogEntry{},
	}

	if err := u.vehiclePersistence.CreateVehicle(ctx, vehicle); err != nil {
		return nil, err
	}
	return &CreateVehicleResponse{Vehicle: vehicle}, nil
}
